# DapurPintar AI Authentication and Authorization Architecture

## Document Control

| Item | Value |
|---|---|
| Status | Proposed for M2 review |
| Scope | MVP identity, authentication, authorization, and security boundaries |
| Primary audience | Product, backend, frontend, AI, QA, DevOps, and AI-assisted development |
| Related diagram | `docs/architecture/diagrams/authentication-authorization-flow.mmd` |

## Executive Summary

This document defines how DapurPintar AI establishes identity, authenticates account participation, authorizes access to user-owned business resources, and protects security-sensitive operations in the MVP. It translates the Architecture Vision, Database Design, and API Design into security boundaries without defining Go handlers, middleware code, SQL migrations, or deployment configuration.

The MVP uses secure password authentication with short-lived JWT access tokens and rotating refresh sessions. The browser receives session credentials through secure, HttpOnly, SameSite cookies. Durable account and session-revocation facts remain authoritative in PostgreSQL. Redis may accelerate rate limiting and bounded session coordination, but loss of Redis must not grant access or remove an account restriction.

Authentication answers **who is acting**. Authorization answers **what that identity may access or change**. A valid account identity never grants access to another user's Profile, Pantry, Favorites, Meal Plans, Shopping Lists, Recommendations, or Conversations. Every protected operation is authorized against the authenticated User and the owning bounded context.

## Goals

- Establish trusted account identity for the MVP.
- Protect all user-owned resources from cross-user access.
- Provide predictable registration, login, logout, and current-account behavior.
- Make token compromise, account restriction, logout, and revocation bounded and observable.
- Preserve a clear seam for future household membership and role-based access.
- Keep authentication concerns separate from kitchen business rules.
- Minimize personal data in tokens, logs, telemetry, and provider requests.

## Non-Goals

The following are outside the MVP implementation scope:

- Household membership and shared resource authorization.
- Nutrition Professional, Grocery Partner, or Commercial Operator roles.
- Social login, enterprise SSO, and external identity federation.
- Multi-factor authentication, unless security validation promotes it before launch.
- Passwordless authentication.
- Fine-grained administrative consoles.
- Authorization through client-supplied ownership identifiers.

These capabilities may be added through explicit security and product decisions without weakening the MVP ownership model.

## Security Principles

- Authentication establishes identity; it does not establish permission for every resource.
- Authorization is enforced on every protected read and write.
- The server derives ownership scope from authenticated identity and server-side relationships.
- Client-supplied `userId`, account scope, or ownership fields are never trusted for authorization.
- PostgreSQL is authoritative for Account state and durable session or revocation facts.
- Redis is not an authorization source of truth.
- Tokens contain the minimum claims needed for request authentication and must not contain pantry, preference, recipe, conversation, or AI context.
- Passwords, raw tokens, refresh secrets, provider credentials, and sensitive private context are never logged.
- Security failures return safe responses and must not reveal whether another user's private resource exists.
- Authentication and authorization events are observable without exposing secrets or unnecessary personal data.

## Identity Model

### Account

Identity and Access owns the Account aggregate. The Account establishes trusted participation and provides the security scope used by personal product contexts.

The MVP Account lifecycle is:

```text
Pending -> Active -> Restricted -> Closed
```

- **Pending:** registration has been accepted but the account is not fully active under the product's verification policy.
- **Active:** the account may authenticate and use authorized MVP capabilities.
- **Restricted:** authentication or selected operations are blocked by a security or policy decision.
- **Closed:** the account can no longer participate; retention and deletion behavior follow the approved privacy policy.

The exact email-verification requirement remains a product decision. The authentication boundary must support a Pending state without treating an unverified or restricted account as Active.

### User Profile and ownership scope

One active Account has one User Profile in the MVP. The authenticated Account identifies the User Profile that owns or contextualizes:

- Pantry and Pantry Items.
- Favorites.
- Meal Plans and Planned Meals.
- Shopping Lists and Shopping Items.
- Kitchen Recommendations and the recommendation-scoped Conversation.

Identity and Access owns the trusted identity. User Context and Preferences owns profile meaning. Other contexts own their respective business rules and data.

## Authentication Model

### Credential authentication

Registration and login use a password-based credential flow:

1. The client submits the minimum account credentials over an encrypted connection.
2. The application validates input and account policy.
3. The password is processed with a memory-hard, adaptive password hashing algorithm and is never persisted in plaintext.
4. Login compares the supplied password with the stored password hash using a timing-safe verification operation.
5. Only an Active account can establish normal authenticated participation.
6. Successful authentication creates a refresh session and issues an access token.

Password policy, breach screening, verification, recovery, and lockout thresholds must be finalized before implementation. They must not be implemented as domain rules inside unrelated kitchen modules.

### Access token

The access token is a short-lived signed JWT used to authenticate requests during a bounded session. It contains only security claims such as:

- Issuer and audience.
- Subject identifying the Account or User.
- Issued-at and expiration timestamps.
- Token identifier where revocation or audit correlation requires it.
- Authentication context needed by the approved session policy.

The token does not contain roles or household permissions that the server cannot safely keep current. Authorization remains a server-side policy decision using current Account state, resource ownership, and future scope membership.

### Refresh session

The refresh session is a long-lived, revocable session credential. The server stores only a protected representation of the refresh secret and session metadata required for revocation, expiry, audit, and device/session management.

Refresh rotation is mandatory:

- A valid refresh operation invalidates the previous refresh secret.
- A replacement refresh secret is issued for the same session lineage.
- Reuse of an invalidated refresh secret is treated as suspected token theft.
- Suspected reuse revokes the affected session lineage and produces a security event.
- Logout revokes the current refresh session and clears browser credentials.

Redis may cache a bounded revocation decision, but PostgreSQL remains authoritative when the cache is unavailable or conflicting.

### Browser transport

The primary MVP Web Application uses secure cookies for session credentials:

- `HttpOnly` prevents application JavaScript from reading the credential.
- `Secure` requires encrypted transport.
- `SameSite` is configured to limit cross-site credential sending according to the deployment topology.
- Cookie scope is limited to the required application domain and path.
- State-changing cookie-authenticated requests require CSRF protection appropriate to the final frontend and deployment topology.

The API contract remains transport-aware but does not expose refresh secrets or provider credentials. Any future non-browser consumer requires an explicit authentication profile rather than weakening browser protections.

## Authentication Operations

The public API contract defines these account operations:

| Operation | Access | Security behavior |
|---|---|---|
| `POST /api/v1/accounts` | Public | Creates an account subject to registration policy and abuse controls. |
| `POST /api/v1/accounts/login` | Public | Verifies credentials and starts authenticated participation. |
| `POST /api/v1/accounts/refresh` | Refresh-session authenticated | Rotates a valid refresh session and issues a new access session. |
| `POST /api/v1/accounts/logout` | Authenticated | Revokes the current refresh session and clears session cookies. |
| `GET /api/v1/accounts/me` | Authenticated | Returns the current account participation context, not credentials or secrets. |

Authentication failures use safe, stable error categories. Login must not disclose whether an email or other identifier exists. Registration must not allow account enumeration through distinguishable responses where the product policy requires concealment.

## Authorization Model

### Authorization decision

Every protected request passes through these conceptual checks:

1. The request has valid transport and contract syntax.
2. Authentication establishes a valid, non-expired identity.
3. The Account is in a state allowed to perform the operation.
4. The requested resource is resolved within the authenticated ownership scope.
5. The owning bounded context authorizes the business operation and lifecycle transition.
6. The application executes the use case without accepting client-defined ownership.

Authentication middleware may establish identity, but the Application Layer remains responsible for resource and business authorization. Domain rules remain authoritative for state transitions and invariants.

### MVP ownership policy

The MVP has one personal User scope. The default policy is:

| Resource | Required scope | Owning context |
|---|---|---|
| Account participation | Current Account | Identity and Access |
| User Profile | Current Account's User Profile | User Context and Preferences |
| Pantry and Pantry Items | Current User Profile | Pantry Management |
| Favorites | Current User Profile | Culinary Knowledge and Recipe Experience |
| Meal Plans and Planned Meals | Current User Profile | Meal Planning |
| Shopping Lists and Shopping Items | Current User Profile | Shopping Optimization |
| Recommendations and options | Current User Profile | AI-Assisted Kitchen Decision Support |
| Recommendation Conversation | Owning Recommendation within Current User Profile | AI-Assisted Kitchen Decision Support |

Public recipe discovery is read-only and must return only general recipe information. It must never infer or return personal suitability, pantry availability, favorites, recommendations, or activity.

### Future scope policy

Household Collaboration may introduce a shared Household scope. It must define membership, invitation, role, resource ownership, revocation, and personal-versus-shared data rules before any household endpoint becomes authenticated behavior.

Future roles are capabilities, not automatic access:

- Household Member may access only explicitly shared household resources.
- Nutrition Professional may access approved nutrition capabilities under explicit consent and scope.
- Grocery Partner may access approved purchase-intent capabilities only.
- Commercial Operator may access SaaS commercial administration, not private kitchen content by default.

No future role may bypass the owning bounded context or turn a reference into ownership.

## Authorization Failure Semantics

- **401 Unauthorized:** authentication is missing, malformed, expired, or invalid.
- **403 Forbidden:** the identity is valid but the operation is outside its authorized scope or the Account is restricted.
- **404 Not Found:** a resource is not available within the authorized scope; this may conceal another user's resource existence.
- **409 Conflict:** the request conflicts with current session, account, or aggregate state.
- **422 Unprocessable Entity:** the request violates a recognized input or business rule.

Responses use the API error contract and include a safe error code and request identifier. They do not expose password verification details, token state, ownership identifiers, SQL, stack traces, or internal authorization policy names.

## Abuse Protection

Rate limiting and abuse controls apply to registration, login, refresh, logout, and AI-intensive endpoints. Controls should consider Account, network, device or session signals, and endpoint sensitivity without creating an unsafe privacy footprint.

Security controls include:

- Bounded login attempts and progressive protection against credential guessing.
- Rate limits on registration and refresh attempts.
- Generic authentication failure messages.
- Refresh-token reuse detection.
- Account restriction and security-event escalation.
- Request correlation for investigation without logging secrets.
- Provider and AI cost controls remain separate from identity authorization.

Rate-limit state may use Redis, but a Redis failure must fail closed for security-critical decisions or degrade to a safe bounded policy rather than granting access.

## Audit and Observability

The system records security-relevant events with a request and trace reference where available:

- Account registration accepted or rejected.
- Login success and failure category.
- Refresh rotation and suspected refresh reuse.
- Logout and session revocation.
- Account restriction or closure.
- Authorization denial for sensitive operations.
- Password or credential policy changes when supported.
- Future privileged role or household membership changes.

Audit records must use safe identifiers, timestamps, outcome, reason category, and source context. They must exclude passwords, raw tokens, full cookies, provider secrets, full prompts, and unnecessary private kitchen data.

Operational metrics should cover authentication latency, success and failure rates, refresh failures, revocation events, rate-limit outcomes, authorization denials, and unusual usage patterns. OpenTelemetry instrumentation follows the existing observability architecture and applies attribute redaction.

## Data Protection and Lifecycle

- Password hashes are protected as credential secrets and are never returned by the API.
- Refresh session records are protected, expire according to policy, and are revocable.
- Closed-account handling follows the product's retention, deletion, and export policy.
- Access tokens expire without requiring a database write for normal requests.
- Session revocation and Account restriction take precedence over cached authentication state.
- Secrets are supplied through deployment secret management and are rotated without exposing them in source control or logs.
- Authentication data is not included in AI context assembly.

## Threats and Mitigations

| Threat | Architectural mitigation |
|---|---|
| Password guessing | Adaptive password hashing, rate limits, generic failures, and account protection. |
| Stolen access token | Short expiry, encrypted transport, minimal claims, and Account/session checks. |
| Stolen refresh token | Rotation, protected storage, reuse detection, lineage revocation, and audit events. |
| Cross-user data access | Server-derived ownership scope and owning-context authorization. |
| Session fixation | New session lineage after authentication and refresh rotation. |
| CSRF | Secure SameSite cookies plus an explicit CSRF defense for state-changing requests. |
| XSS credential theft | HttpOnly cookies, output encoding, input validation, and frontend security policy. |
| User enumeration | Safe, non-distinguishing authentication and registration responses. |
| Redis failure | PostgreSQL authority and fail-closed behavior for security-critical checks. |
| Logging leakage | Redaction of credentials, tokens, private context, and provider data. |
| Privilege escalation | Explicit future role policies and no client-controlled ownership or role claims. |

## Decisions Deferred for Validation

The following details require product, security, or deployment validation before implementation:

- Email verification requirement and verification lifecycle.
- Password recovery and credential-change flows.
- Exact access-token and refresh-session lifetimes.
- Final cookie domain, SameSite mode, and CSRF mechanism.
- Multi-factor authentication requirement before public launch.
- Account restriction thresholds and manual recovery process.
- Retention and deletion periods for account and session audit data.
- Non-browser authentication profiles for future partners or public APIs.

These decisions must preserve the principles and ownership boundaries in this document.

## Risks and Assumptions

### Risks

- Cookie and frontend deployment topology may require a different CSRF or session transport profile.
- JWT claims may become stale if authorization is incorrectly inferred from token contents.
- Refresh-session storage or rotation failures may create unexpected user sign-outs or security exposure.
- Rate limits may block legitimate users if they are not calibrated with product usage and regional network behavior.
- Future household sharing may expose personal context if shared and private ownership is not explicitly separated.
- Audit records may become a source of sensitive-data leakage if redaction is not enforced.

### Assumptions

- The MVP has one personal User scope and no household authorization runtime.
- PostgreSQL remains the authoritative store for Account and durable session or revocation state.
- Redis is supporting infrastructure only.
- The Web Application is the primary MVP client.
- OpenAI and other AI providers never receive authentication credentials or raw session data.
- `/api/v1` remains the public API prefix.
- Authentication implementation will use secure password handling and JWT-based access control as established by Architecture Vision.
- Authorization decisions remain in the application and owning bounded contexts, not in the frontend.

## Exit Criteria

M2-012 is ready for review when:

- Identity and Account lifecycle boundaries are defined.
- Authentication operations align with the API Design.
- Token, session, logout, and revocation behavior are defined at architecture level.
- MVP ownership authorization is explicit for every protected resource.
- Future roles and household scope are isolated from MVP behavior.
- Abuse protection, auditability, privacy, and failure behavior are covered.
- The flow diagram reflects the authentication and authorization boundaries.

## Related Documents

- `docs/architecture/architecture-vision.md`
- `docs/architecture/api-design.md`
- `docs/architecture/database-design.md`
- `docs/architecture/bounded-context.md`
- `docs/architecture/adr/ADR-007-use-rest-api.md`
- `docs/architecture/adr/ADR-003-use-postgresql-as-system-of-record.md`
- `docs/architecture/diagrams/authentication-authorization-flow.mmd`
