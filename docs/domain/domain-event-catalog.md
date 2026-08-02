# DapurPintar AI Domain Event Catalog

## Purpose

This catalog is the formal business reference for domain events identified by the DapurPintar AI Event Storming work. It defines the meaning, ownership, lifecycle context, and business relationships of each event.

This catalog does not define message formats, transport, storage, API contracts, or implementation mechanisms. The attributes below are business attributes required to understand the event, not a database or integration schema.

## Event Status

| Status | Meaning |
|---|---|
| Confirmed | Accepted as a meaningful MVP business fact. |
| Candidate | Identified during discovery but requires Product and domain validation. |
| Future | Reserved for a post-MVP business capability. |

## Catalog Conventions

- Every event is written in the past tense because it represents a business fact that has occurred.
- The Owning Context is the authority for the event's meaning and lifecycle.
- The Trigger is the command or business policy that may cause the event.
- Consumers use the event's business meaning but do not become event owners.
- Follow-up Events are possible business consequences, not guaranteed automatic actions.
- An event must not claim facts owned by another Aggregate or Bounded Context.
- Event attributes are purpose-limited and must respect personal and household privacy.

## MVP Events

### Account Registered

- **Status:** Confirmed
- **Owner:** Identity and Access
- **Triggered by:** Register Account
- **Aggregate:** Account
- **Business Meaning:** A new trusted account has been established and may begin participation in the product.
- **Required Attributes:** Account identity; registration context; participation status; occurred at.
- **Consumers:** User Context and Preferences; product onboarding.
- **Business Rules:** Account identity must satisfy registration requirements; the account must not be active before required registration conditions are met.
- **Follow-up Events:** User Profile Created, when profile setup begins or completes; Account Access Changed, when participation status changes.

### Account Access Changed

- **Status:** Confirmed
- **Owner:** Identity and Access
- **Triggered by:** Confirm Account Participation; Restrict Account; Close Account
- **Aggregate:** Account
- **Business Meaning:** The account's ability to participate in protected kitchen decisions has changed.
- **Required Attributes:** Account identity; previous access state; new access state; reason; occurred at.
- **Consumers:** User Context and Preferences; protected kitchen contexts; future Commercial Participation.
- **Business Rules:** A restricted or closed account cannot act as an active participant; access state does not change the meaning of kitchen facts.
- **Follow-up Events:** None required; downstream contexts may reassess participation.

### User Profile Created

- **Status:** Candidate
- **Owner:** User Context and Preferences
- **Triggered by:** Create User Profile
- **Aggregate:** User Profile
- **Business Meaning:** Personal cooking context has been initialized for a trusted account.
- **Required Attributes:** Account identity; profile identity; initial context status; occurred at.
- **Consumers:** AI-Assisted Kitchen Decision Support; onboarding; future Family and Household Collaboration.
- **Business Rules:** The profile must belong to a valid account; incomplete context must not be presented as complete personalization context.
- **Follow-up Events:** User Preferences Changed, when the user supplies or changes preferences.

### User Preferences Changed

- **Status:** Confirmed
- **Owner:** User Context and Preferences
- **Triggered by:** Update Preferences; Update Cooking Constraints; Update Cooking Goals
- **Aggregate:** User Profile
- **Business Meaning:** The personal cooking context used to guide future decisions has changed.
- **Required Attributes:** Profile identity; changed preference or constraint meaning; effective context; changed by; occurred at.
- **Consumers:** AI-Assisted Kitchen Decision Support; future Nutrition Guidance; future Family and Household Collaboration where sharing is explicit.
- **Business Rules:** The change must be attributable to the profile owner; personal context is private by default; invalid or contradictory context must be rejected or clarified.
- **Follow-up Events:** Kitchen Recommendation Requested, only when a product flow explicitly requests new guidance.

### Pantry Item Added

- **Status:** Confirmed
- **Owner:** Pantry Management
- **Triggered by:** Add Pantry Item
- **Aggregate:** Pantry
- **Business Meaning:** A new ingredient has become part of the user's pantry context.
- **Required Attributes:** Pantry identity; ingredient identity; quantity; unit; expiry context when known; added by; occurred at.
- **Consumers:** AI-Assisted Kitchen Decision Support; Pantry Dashboard; Shopping Optimization.
- **Business Rules:** Ingredient identity must be valid; quantity must be greater than zero for an available item; the item must belong to the correct pantry scope.
- **Follow-up Events:** Ingredient Approaching Expiry, when expiry policy is met; Kitchen Recommendation Requested, when recommendation policy identifies useful new context.

### Pantry Item Adjusted

- **Status:** Confirmed
- **Owner:** Pantry Management
- **Triggered by:** Adjust Pantry Item
- **Aggregate:** Pantry
- **Business Meaning:** The quantity or expiry context of an existing pantry ingredient has changed.
- **Required Attributes:** Pantry identity; pantry item identity; previous quantity or expiry context; new quantity or expiry context; adjusted by; occurred at.
- **Consumers:** AI-Assisted Kitchen Decision Support; Shopping Optimization; Pantry Dashboard.
- **Business Rules:** Quantity cannot become negative; the item must belong to the pantry being adjusted; the adjustment must preserve the pantry's ingredient meaning.
- **Follow-up Events:** Ingredient Approaching Expiry; Kitchen Recommendation Requested, when policy determines that the changed context matters.

### Pantry Item Removed

- **Status:** Confirmed
- **Owner:** Pantry Management
- **Triggered by:** Remove Pantry Item
- **Aggregate:** Pantry
- **Business Meaning:** An ingredient is no longer considered part of the user's pantry context.
- **Required Attributes:** Pantry identity; pantry item identity; ingredient identity; removal reason; removed by; occurred at.
- **Consumers:** AI-Assisted Kitchen Decision Support; Shopping Optimization; Pantry Dashboard.
- **Business Rules:** Removal must not silently create a purchase intention; active recommendations and plans may need review if they relied on the removed ingredient.
- **Follow-up Events:** Kitchen Recommendation Requested or Planned Meal Changed, only when an explicit business decision is made.

### Ingredient Approaching Expiry

- **Status:** Confirmed
- **Owner:** Pantry Management
- **Triggered by:** Review Pantry Expiry; Expiry Policy
- **Aggregate:** Pantry
- **Business Meaning:** A pantry ingredient requires attention because its expiry context meets the agreed threshold.
- **Required Attributes:** Pantry identity; pantry item identity; ingredient identity; expiry context; attention threshold; occurred at.
- **Consumers:** AI-Assisted Kitchen Decision Support; Pantry Dashboard; future Notifications and Reminders.
- **Business Rules:** The event must be based on known expiry context; it must not claim an ingredient is expired when the business threshold only indicates approaching expiry.
- **Follow-up Events:** Kitchen Recommendation Requested; future Reminder Requested.

### Recipe Favorited

- **Status:** Confirmed
- **Owner:** Culinary Knowledge and Recipe Experience
- **Triggered by:** Favorite Recipe
- **Aggregate:** Recipe
- **Business Meaning:** A user has expressed an enduring preference for a recipe.
- **Required Attributes:** User identity; recipe identity; favorite status; occurred at.
- **Consumers:** Recipe Discovery; AI-Assisted Kitchen Decision Support; future Recommendation Feedback.
- **Business Rules:** The recipe must exist and be available to the user; a favorite is a preference signal, not proof of current suitability.
- **Follow-up Events:** None required; future personalization may use the signal under an explicit policy.

### Kitchen Recommendation Requested

- **Status:** Confirmed
- **Owner:** AI-Assisted Kitchen Decision Support
- **Triggered by:** Request Kitchen Recommendation; Use-First Recommendation; Recommendation Policy
- **Aggregate:** Kitchen Recommendation
- **Business Meaning:** A user or product flow has requested contextual cooking decision support.
- **Required Attributes:** Requesting user or product context; decision purpose; relevant context scope; requested at.
- **Consumers:** AI-Assisted Kitchen Decision Support; External AI Provider as a future or current participant.
- **Business Rules:** The request must have an identifiable user context and a legitimate decision purpose; personal context must be minimized.
- **Follow-up Events:** Kitchen Recommendation Created; Kitchen Recommendation Unable to Complete, candidate.

### Kitchen Recommendation Created

- **Status:** Confirmed
- **Owner:** AI-Assisted Kitchen Decision Support
- **Triggered by:** Kitchen Recommendation Requested; Recommendation Suitability Policy
- **Aggregate:** Kitchen Recommendation
- **Business Meaning:** A contextual recommendation has been produced with one or more cooking options and a relevance basis or limitation.
- **Required Attributes:** Recommendation identity; user context identity; recommendation options; rationale; context limitations; created at.
- **Consumers:** Recommendation View; Meal Planning; Shopping Optimization; User.
- **Business Rules:** The recommendation must not claim facts owned by Pantry, Recipe, Meal Plan, Shopping, or future Nutrition contexts; unsupported output must not be presented as trusted guidance.
- **Follow-up Events:** Recommendation Presented; Recommendation Rejected; Recommendation Superseded.

### Recommendation Presented

- **Status:** Candidate
- **Owner:** AI-Assisted Kitchen Decision Support
- **Triggered by:** Present Recommendation
- **Aggregate:** Kitchen Recommendation
- **Business Meaning:** A created recommendation has been made available for the user to consider.
- **Required Attributes:** Recommendation identity; presentation context; user identity; presented at.
- **Consumers:** User; Recommendation View; future Recommendation Feedback.
- **Business Rules:** Only a recommendation with a meaningful rationale or explicit limitation may be presented; presentation is not acceptance.
- **Follow-up Events:** Recommendation Accepted; Recommendation Rejected; Recommendation Superseded.

### Recommendation Accepted

- **Status:** Confirmed
- **Owner:** AI-Assisted Kitchen Decision Support
- **Triggered by:** Accept Recommendation
- **Aggregate:** Kitchen Recommendation
- **Business Meaning:** The user has accepted a recommendation as useful for further action.
- **Required Attributes:** Recommendation identity; accepted option; accepting user; acceptance context; occurred at.
- **Consumers:** Meal Planning; Shopping Optimization; future Recommendation Feedback.
- **Business Rules:** Only a presented recommendation can be accepted; acceptance does not automatically create a meal plan or complete a purchase; an accepted recommendation cannot return to Draft or Proposed state.
- **Follow-up Events:** Meal Plan Created or Meal Planned, only after an explicit planning command; Shopping List Generated, only after an explicit generation decision.

### Recommendation Rejected

- **Status:** Confirmed
- **Owner:** AI-Assisted Kitchen Decision Support
- **Triggered by:** Reject Recommendation
- **Aggregate:** Kitchen Recommendation
- **Business Meaning:** The user has declined a recommendation for the current decision.
- **Required Attributes:** Recommendation identity; rejecting user; rejection reason when provided; occurred at.
- **Consumers:** AI-Assisted Kitchen Decision Support; future Recommendation Feedback.
- **Business Rules:** Rejection does not prove the recommendation was factually incorrect; it must not silently change the user's explicit preferences.
- **Follow-up Events:** Kitchen Recommendation Requested, only when the user or policy requests another option.

### Recommendation Superseded

- **Status:** Candidate
- **Owner:** AI-Assisted Kitchen Decision Support
- **Triggered by:** Replace Recommendation; Context Changed
- **Aggregate:** Kitchen Recommendation
- **Business Meaning:** A previous recommendation is no longer the current guidance because a new option or changed context replaced it.
- **Required Attributes:** Previous recommendation identity; replacement context or recommendation identity; reason; occurred at.
- **Consumers:** Recommendation View; Meal Planning; Shopping Optimization.
- **Business Rules:** A superseded recommendation must not be presented as current guidance; any accepted downstream commitment requires its own review.
- **Follow-up Events:** Kitchen Recommendation Created for the replacement guidance.

### Meal Plan Created

- **Status:** Confirmed
- **Owner:** Meal Planning
- **Triggered by:** Create Meal Plan
- **Aggregate:** Meal Plan
- **Business Meaning:** A daily or weekly planning intention has been established.
- **Required Attributes:** Meal plan identity; owner identity; planning period; creation context; occurred at.
- **Consumers:** User; Meal Planning; Shopping Optimization; Weekly Meal View.
- **Business Rules:** The planning period must be defined; the plan must belong to the correct user or future household scope.
- **Follow-up Events:** Meal Planned; Planned Meal Changed; Meal Plan Completed, candidate.

### Meal Planned

- **Status:** Confirmed
- **Owner:** Meal Planning
- **Triggered by:** Plan Meal
- **Aggregate:** Meal Plan
- **Business Meaning:** A meal intention has been assigned to a date or meal occasion within a planning period.
- **Required Attributes:** Meal plan identity; planned meal identity; recipe or cooking option identity; meal date; meal slot; planning context; occurred at.
- **Consumers:** Shopping Optimization; Weekly Meal View; AI-Assisted Kitchen Decision Support.
- **Business Rules:** One meal slot cannot contain two active meals where overlap is not permitted; a planned meal must belong to one planning period; accepted guidance remains distinct from the plan commitment.
- **Follow-up Events:** Shopping List Generated, when purchase needs are identified; Planned Meal Changed.

### Planned Meal Changed

- **Status:** Confirmed
- **Owner:** Meal Planning
- **Triggered by:** Change Planned Meal; Remove Planned Meal
- **Aggregate:** Meal Plan
- **Business Meaning:** A previously intended meal or its place in the plan has changed.
- **Required Attributes:** Meal plan identity; planned meal identity; previous intention; new intention or removal reason; changed by; occurred at.
- **Consumers:** Shopping Optimization; AI-Assisted Kitchen Decision Support; Weekly Meal View.
- **Business Rules:** The change must preserve planning-period and meal-slot invariants; related shopping intent must be reviewed rather than silently assumed to change.
- **Follow-up Events:** Shopping List Generated or revised shopping intent, after explicit decision.

### Meal Plan Completed

- **Status:** Candidate
- **Owner:** Meal Planning
- **Triggered by:** Complete Meal Plan
- **Aggregate:** Meal Plan
- **Business Meaning:** The planning period or meal intention has been fulfilled according to the agreed product meaning.
- **Required Attributes:** Meal plan identity; completed period; completion status; completed by or completion reason; occurred at.
- **Consumers:** Weekly Meal View; Success Measurement; future Recommendation Feedback.
- **Business Rules:** Completion does not claim that every ingredient was purchased or every meal was cooked unless the business definition explicitly says so.
- **Follow-up Events:** None required.

### Shopping List Generated

- **Status:** Confirmed
- **Owner:** Shopping Optimization
- **Triggered by:** Generate Shopping List; Shopping Need Assessment
- **Aggregate:** Shopping List
- **Business Meaning:** Purchase intent has been created from meal intent, pantry gaps, or an explicit user request.
- **Required Attributes:** Shopping list identity; owner identity; source meal or cooking intent; ingredient needs; generation context; occurred at.
- **Consumers:** Shopping Review View; User; future Grocery Partner.
- **Business Rules:** Available pantry quantities should be considered; generated items must remain reviewable; generation does not equal purchase completion.
- **Follow-up Events:** Shopping Item Completed; Shopping List Completed, candidate; Shopping List Revised, candidate.

### Shopping Item Completed

- **Status:** Confirmed
- **Owner:** Shopping Optimization
- **Triggered by:** Complete Shopping Item
- **Aggregate:** Shopping List
- **Business Meaning:** One purchase intention has been fulfilled or closed by the user.
- **Required Attributes:** Shopping list identity; shopping item identity; ingredient identity; completion status; completed by; occurred at.
- **Consumers:** Shopping View; future Marketplace or Grocery Partner; Pantry Management only as a prompt for explicit user action.
- **Business Rules:** A completed item cannot remain open; completion does not automatically add the ingredient to the Pantry.
- **Follow-up Events:** Pantry Item Added only after an explicit Add Pantry Item command.

### Shopping List Completed

- **Status:** Candidate
- **Owner:** Shopping Optimization
- **Triggered by:** Complete Shopping List
- **Aggregate:** Shopping List
- **Business Meaning:** The user's shopping intention for a list is closed.
- **Required Attributes:** Shopping list identity; completion status; completed by; remaining item context; occurred at.
- **Consumers:** Shopping View; future Marketplace; Success Measurement.
- **Business Rules:** The meaning of completion must clarify whether unresolved items are allowed; completion does not claim that pantry availability changed.
- **Follow-up Events:** None required.

## Future Events

### Household Collaboration Started

- **Status:** Future
- **Owner:** Family and Household Collaboration
- **Triggered by:** Start Household Collaboration
- **Aggregate:** Household Collaboration
- **Business Meaning:** A shared household context for kitchen decisions has been established.
- **Required Attributes:** Household identity; initiating person; participation purpose; visibility scope; occurred at.
- **Consumers:** Pantry Management; Meal Planning; Shopping Optimization; User Context and Preferences.
- **Business Rules:** Shared participation must be explicit; personal preferences remain private unless shared under an agreed rule.
- **Follow-up Events:** Household Member Participation Changed.

### Household Member Participation Changed

- **Status:** Future
- **Owner:** Family and Household Collaboration
- **Triggered by:** Invite Household Member; Assign Participation; Remove Household Member
- **Aggregate:** Household Collaboration
- **Business Meaning:** A person's ability to participate in shared kitchen decisions has changed.
- **Required Attributes:** Household identity; person identity; previous participation; new participation; changed by; occurred at.
- **Consumers:** Pantry Management; Meal Planning; Shopping Optimization.
- **Business Rules:** Participation scope must be explicit; household access must not expose unrelated personal context.
- **Follow-up Events:** None required.

### Nutrition Goal Changed

- **Status:** Future
- **Owner:** Nutrition Guidance
- **Triggered by:** Set Nutrition Goal; Update Nutrition Goal
- **Aggregate:** Nutrition Goal
- **Business Meaning:** A nutrition-related goal or constraint has been established or revised.
- **Required Attributes:** Person or household scope; nutrition goal; effective context; changed by; occurred at.
- **Consumers:** AI-Assisted Kitchen Decision Support; Meal Planning; Recipe Experience.
- **Business Rules:** Nutrition meaning must come from the Nutrition Guidance context; the event must not be interpreted as medical advice without validated policy.
- **Follow-up Events:** Nutrition-Aware Recommendation Created, candidate.

### Reminder Requested

- **Status:** Future
- **Owner:** Notifications and Reminders
- **Triggered by:** Ingredient Approaching Expiry; Meal Planned; Kitchen Recommendation policy
- **Aggregate:** Notification
- **Business Meaning:** A business situation requires a timely user prompt.
- **Required Attributes:** Recipient scope; reason; related business context; desired timing; occurred at.
- **Consumers:** User; Notifications and Reminders.
- **Business Rules:** The originating context owns why the reminder matters; reminder delivery must respect user preferences and privacy.
- **Follow-up Events:** Reminder Delivered; Reminder Dismissed, candidate.

### Commercial Participation Changed

- **Status:** Future
- **Owner:** SaaS Commercial Operations
- **Triggered by:** Start Subscription; Change Participation; Cancel Subscription
- **Aggregate:** Commercial Participation
- **Business Meaning:** The customer's access or participation level in commercial product capabilities has changed.
- **Required Attributes:** Customer identity; previous participation; new participation; effective period; reason; occurred at.
- **Consumers:** AI-Assisted Kitchen Decision Support; User Profile; future premium capabilities.
- **Business Rules:** Commercial participation may constrain access or usage but must not change kitchen concept meaning or recommendation suitability.
- **Follow-up Events:** Premium Access Granted or Premium Access Ended, candidate.

## Event Catalog Maintenance

Event definitions should be reviewed when:

- A bounded context gains or loses ownership.
- An event's business meaning changes.
- A policy changes the conditions under which the event occurs.
- A lifecycle transition is added, removed, or renamed.
- A new consumer needs a different business interpretation.
- MVP validation reveals that an event is only a user interface activity, not a domain fact.

The catalog should remain synchronized with Event Storming, the bounded-context map, the tactical DDD model, and future domain decisions. Changes should be discussed with Product and domain stakeholders before being treated as stable business language.
