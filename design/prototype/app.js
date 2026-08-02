const screens = [...document.querySelectorAll('.screen')];
const navItems = [...document.querySelectorAll('[data-screen]')];
const toast = document.getElementById('toast');
const sidebar = document.querySelector('.sidebar');

function showToast(message) {
  toast.textContent = message;
  toast.classList.add('show');
  window.clearTimeout(showToast.timeout);
  showToast.timeout = window.setTimeout(() => toast.classList.remove('show'), 2600);
}

function showScreen(name) {
  screens.forEach((screen) => screen.classList.toggle('active', screen.id === `screen-${name}`));
  document.querySelectorAll('.nav-item, .mobile-nav button').forEach((item) => item.classList.toggle('active', item.dataset.screen === name));
  sidebar.classList.remove('open');
  window.scrollTo({ top: 0, behavior: 'smooth' });
}

navItems.forEach((item) => item.addEventListener('click', (event) => {
  event.preventDefault();
  showScreen(item.dataset.screen);
}));

document.getElementById('mobileMenu').addEventListener('click', () => sidebar.classList.toggle('open'));

const pantryModal = document.getElementById('pantryModal');
const pantryForm = document.getElementById('pantryForm');
const ingredientName = document.getElementById('ingredientName');

function openPantryForm() {
  pantryModal.showModal();
  window.setTimeout(() => ingredientName.focus(), 50);
}

document.getElementById('openPantryForm').addEventListener('click', openPantryForm);
document.getElementById('openPantryFormEmpty').addEventListener('click', openPantryForm);
pantryForm.addEventListener('submit', (event) => {
  event.preventDefault();
  pantryModal.close();
  document.getElementById('itemCount').textContent = '13';
  showToast(`${ingredientName.value || 'Item'} saved to your pantry.`);
  ingredientName.value = '';
});

function acceptRecommendation() {
  document.getElementById('acceptedNote').hidden = false;
  document.getElementById('acceptOption').textContent = 'Option accepted';
  document.getElementById('acceptOption').disabled = true;
  document.getElementById('recommendationAccepted').hidden = false;
  showToast('Option accepted. Choose your next action.');
}

document.getElementById('acceptOption').addEventListener('click', acceptRecommendation);
document.querySelectorAll('.option-accept').forEach((button) => button.addEventListener('click', () => {
  document.querySelectorAll('.option-card').forEach((card) => card.classList.remove('selected'));
  button.closest('.option-card').classList.add('selected');
  acceptRecommendation();
}));

document.getElementById('conversationButton').addEventListener('click', () => {
  showToast('Conversation is scoped to this recommendation.');
});

document.getElementById('toggleAi').addEventListener('click', (event) => {
  const card = document.getElementById('recommendationCard');
  const unavailable = card.classList.toggle('ai-unavailable');
  event.currentTarget.textContent = unavailable ? 'Restore AI suggestion' : 'Simulate AI unavailable';
  card.querySelector('.reason').textContent = unavailable ? 'AI could not create a suggestion right now. Your pantry is still available.' : 'Uses spinach first and fits your time preference.';
  card.querySelector('.card-actions').innerHTML = unavailable
    ? '<button class="button button-primary" data-screen="discover">Browse recipes</button><button class="button button-quiet" id="retryAi">Try again</button>'
    : '<button class="button button-primary" id="acceptOption">Use this idea</button><button class="button button-quiet" data-screen="recommendation">Ask about it</button>';
  card.querySelectorAll('[data-screen]').forEach((item) => item.addEventListener('click', () => showScreen(item.dataset.screen)));
  const accept = document.getElementById('acceptOption');
  if (accept) accept.addEventListener('click', acceptRecommendation);
  const retry = document.getElementById('retryAi');
  if (retry) retry.addEventListener('click', () => event.currentTarget.click());
});

document.getElementById('activateList').addEventListener('click', () => {
  document.getElementById('shoppingStatus').textContent = 'Active';
  document.getElementById('activateList').textContent = 'List active';
  document.getElementById('activateList').disabled = true;
  showToast('Shopping list activated for review in store.');
});

document.querySelectorAll('.shopping-item input').forEach((input) => input.addEventListener('change', () => {
  if (input.checked && !input.closest('.completed')) showToast('Shopping item completed. Pantry was not changed.');
}));

window.addEventListener('hashchange', () => {
  const target = window.location.hash.replace('#', '');
  if (target && document.getElementById(`screen-${target}`)) showScreen(target);
});
