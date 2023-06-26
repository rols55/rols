import { gossips } from './gossip-grid.data.js';

export const grid = () => {
  const container = document.createElement('div');
  container.classList.add('grid');

  const form = createGossipCardForm(); // Create the gossip card form
  container.appendChild(form);

  const rangesDiv = document.createElement('div');
  rangesDiv.classList.add('ranges');

  const widthRange = createRangeInput('width', 200, 800, 200);
  rangesDiv.appendChild(widthRange);

  const fontSizeRange = createRangeInput('fontSize', 20, 40, 20);
  rangesDiv.appendChild(fontSizeRange);

  const backgroundRange = createRangeInput('background', 20, 75, 20);
  rangesDiv.appendChild(backgroundRange);

  container.appendChild(rangesDiv);

  const existingGrid = document.querySelector('.grid');
  if (existingGrid) {
    existingGrid.replaceWith(container);
  } else {
    document.body.appendChild(container);
  }

  gossips.forEach((gossip, index) => {
    const card = createGossipCard(gossip);
    if (index === 0) {
      container.insertBefore(card, form.nextSibling); // Insert the form card after the form
    } else {
      container.appendChild(card);
    }
  });
};

const createGossipCardForm = () => {
  const form = document.createElement('form');
  form.addEventListener('submit', (event) => {
    event.preventDefault();
    const textarea = form.querySelector('textarea');
    const gossip = textarea.value.trim();
    if (gossip) {
      const card = createGossipCard(gossip);
      form.parentNode.insertBefore(card, form.nextSibling); // Insert the new gossip card after the form
      textarea.value = '';
    }
  });

  const textarea = document.createElement('textarea');
  form.appendChild(textarea);

  const submitButton = document.createElement('button');
  submitButton.textContent = 'Share gossip!';
  form.appendChild(submitButton);

  form.classList.add('gossip');

  return form;
};

const createGossipCard = (gossip) => {
  const card = document.createElement('div');
  card.classList.add('gossip');
  card.textContent = gossip;

  return card;
};

const createRangeInput = (id, min, max, defaultValue) => {
  const rangeInput = document.createElement('input');
  rangeInput.type = 'range';
  rangeInput.classList.add('range');
  rangeInput.id = id;
  rangeInput.min = min;
  rangeInput.max = max;
  rangeInput.defaultValue = defaultValue;

  rangeInput.addEventListener('input', () => {
    const value = rangeInput.value;
    console.log(value)
    if (id === 'width') {
      document.querySelectorAll('.gossip').forEach((gossipCard) => {
        gossipCard.style.width = `${value}px`;
      });
    } else if (id === 'fontSize') {
      document.querySelectorAll('.gossip').forEach((gossipCard) => {
        gossipCard.style.fontSize = `${value}px`;
      });
    } else if (id === 'background') {
      let lightness = (value);
      document.querySelectorAll('.gossip').forEach((gossipCard) => {
        gossipCard.style.background =`hsl(240, 100%, ${lightness}%)`      });
    }
  });

  return rangeInput;
};
