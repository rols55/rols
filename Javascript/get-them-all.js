export const getArchitects = () => {
  const allElements = Array.from(document.body.getElementsByTagName('*'));
  const architects = allElements.filter(element => element.tagName.toLowerCase() === 'a');
  const nonArchitects = allElements.filter(element => element.tagName.toLowerCase() !== 'a');
  return [architects, nonArchitects];
};

export const getClassical = () => {
  const allElements = Array.from(document.getElementsByTagName('a'));
  const classicals = allElements.filter(element => element.classList.contains('classical'));
  const nonClassicals = allElements.filter(element => !element.classList.contains('classical'));
  return [classicals, nonClassicals];
};

export const getActive = () => {
  const classicalElements = Array.from(document.getElementsByClassName('classical'));
  const activeClassicals = classicalElements.filter(element => element.classList.contains('active'));
  const nonActiveClassicals = classicalElements.filter(element => !element.classList.contains('active'));
  return [activeClassicals, nonActiveClassicals];
};

export const getBonannoPisano = () => {
  const bonannoPisano = document.getElementById('BonannoPisano');
  const activeClassicals = Array.from(document.querySelectorAll('.classical.active'));
  const otherActiveClassicals = activeClassicals.filter(element => element.id !== 'BonannoPisano');
  return [bonannoPisano, otherActiveClassicals];
};