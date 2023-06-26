import { places } from './where-do-we-go.data.js';

export const explore = () => {
  // Sort places from north to south
  
    places.sort((a, b) => {
      const aLatitude = parseCoordinate(a.coordinates);
      const bLatitude = parseCoordinate(b.coordinates);
      return bLatitude - aLatitude;
    
  })
  
  function parseCoordinate(coordinate) {
    const [degrees, minutes] = coordinate.split('°');
    return parseFloat(degrees) + parseFloat(minutes) / 60;
  }
  

  const sectionContainer = document.createElement('div');
  sectionContainer.classList.add('section-container');

  const locationIndicator = document.createElement('a');
  locationIndicator.classList.add('location');
  let currentPlaceIndex = 0;
  updateLocationIndicator();

  let compassDirection;
  let lastScrollY = window.scrollY;

  window.addEventListener('scroll', () => {
    const scrollY = window.scrollY;
    const middleHeight = window.innerHeight / 2;

    const currentIndex = Math.floor(scrollY / window.innerHeight);
    const nextIndex = currentIndex + 1;

    if (nextIndex < places.length) {
      const nextPlace = places[nextIndex];

      if (scrollY + middleHeight >= nextIndex * window.innerHeight) {
        locationIndicator.textContent = `${nextPlace.name}\n${nextPlace.coordinates}`;
        locationIndicator.style.color = nextPlace.color;

        let compassSymbol = scrollY > lastScrollY ? 'S' : 'N';
        //compass.textContent = compassSymbol;
      }
    }

    lastScrollY = scrollY;
  });

  places.forEach((place) => {
    const section = document.createElement('section');
    const imageURL = `./where-do-we-go_images/${place.name.split(',')[0].replaceAll(' ', '-').toLowerCase()}.jpg`;
    section.style.background = `url(${imageURL})center/cover`;
    sectionContainer.appendChild(section);
  });

  document.body.appendChild(sectionContainer);
  document.body.appendChild(locationIndicator);

  function updateLocationIndicator() {
    const currentPlace = places[currentPlaceIndex];
    locationIndicator.textContent = `${currentPlace.name}\n${currentPlace.coordinates}`;
    locationIndicator.style.color = currentPlace.color;
    locationIndicator.href = `https://www.google.com/maps?q=${currentPlace.coordinates}`;
  }

  function updateCompass() {
    if (!compassDirection) {
      compassDirection = document.createElement('div');
      compassDirection.classList.add('direction');
      document.body.appendChild(compassDirection);
    }
    compassDirection.textContent = window.scrollY > lastScrollY ? 'N' : 'S';
  }
};
