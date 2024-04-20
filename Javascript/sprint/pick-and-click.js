export const pick = () => {
    const body = document.body;
    const hslDiv = document.createElement('div');
    hslDiv.classList.add('hsl');
    const hueDiv = document.createElement('div');
    hueDiv.classList.add('hue');
    hueDiv.style.position = 'absolute';
    hueDiv.style.top = '0';
    hueDiv.style.right = '0';
    const luminosityDiv = document.createElement('div');
    luminosityDiv.classList.add('luminosity');
    luminosityDiv.style.position = 'absolute';
    luminosityDiv.style.bottom = '0';
    luminosityDiv.style.left = '0';
    const axisX = document.createElementNS('http://www.w3.org/2000/svg', 'line');
    axisX.setAttribute('id', 'axisX');
    const axisY = document.createElementNS('http://www.w3.org/2000/svg', 'line');
    axisY.setAttribute('id', 'axisY');
  
    document.addEventListener('mousemove', (event) => {
      const { pageX, pageY } = event;
      const hue = Math.round((pageX / window.innerWidth) * 360);
      const luminosity = Math.round((pageY / window.innerHeight) * 100);
      const hslValue = `hsl(${hue}, 100%, ${luminosity}%)`;
  
      body.style.background = hslValue;
      hslDiv.textContent = hslValue;
      hueDiv.textContent = `Hue: ${hue}`;
      luminosityDiv.textContent = `Luminosity: ${luminosity}%`;
  
      axisX.setAttribute('x1', pageX);
      axisX.setAttribute('x2', pageX);
      axisY.setAttribute('y1', pageY);
      axisY.setAttribute('y2', pageY);
    });
  
    document.addEventListener('click', () => {
      const hslValue = hslDiv.textContent;
      navigator.clipboard.writeText(hslValue);
    });
  
    document.body.appendChild(hslDiv);
    document.body.appendChild(hueDiv);
    document.body.appendChild(luminosityDiv);
    document.body.appendChild(axisX);
    document.body.appendChild(axisY);
  };
  
  // Export the modified pick function
  export default pick;
  