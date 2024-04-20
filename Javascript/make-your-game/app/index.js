document.addEventListener('DOMContentLoaded', () => {
  const grid = document.querySelector('.grid')
  const miniGrid = document.querySelector('.mini-grid')
  
  //Main board builder
  function createBoardByHeight(height) {
    let divs = height * 10
    for (var i = 0; i < divs; i++){
      const newDiv = document.createElement("div");
      grid.appendChild(newDiv);
    }
    for (var i = 0; i < 10; i++){
      const bottom = document.createElement("div");
      bottom.className = "taken";
      grid.appendChild(bottom);
    } 
  }
  createBoardByHeight(20)
  //Mini grid builder
  function CreateMiniGrid(){
    for (var i = 0; i < 16; i++){
      const newMiniGrid = document.createElement("div");
      miniGrid.appendChild(newMiniGrid);
    }
  }
  CreateMiniGrid()

  let squares = Array.from(document.querySelectorAll('.grid div'))
  const scoreDisplay = document.querySelector('#score')
  const startBtn = document.getElementById('start-button')
  var restartbtn = document.getElementById('restart-button');
  const width = 10
  let timerId
  let score = 0
  const colors = [
    'orange',
    'red',
    'purple',
    'green',
    'blue'
  ]
  //The Tetrominoes
  const lTetromino = [
    [1, width+1, width*2+1, 2],
    [width, width+1, width+2, width*2+2],
    [1, width+1, width*2+1, width*2],
    [width, width*2, width*2+1, width*2+2]
  ]
  const zTetromino = [
    [0,width,width+1,width*2+1],
    [width+1, width+2,width*2,width*2+1],
    [0,width,width+1,width*2+1],
    [width+1, width+2,width*2,width*2+1]
  ]
  const tTetromino = [
    [1,width,width+1,width+2],
    [1,width+1,width+2,width*2+1],
    [width,width+1,width+2,width*2+1],
    [1,width,width+1,width*2+1]
  ]
  const oTetromino = [
    [0,1,width,width+1],
    [0,1,width,width+1],
    [0,1,width,width+1],
    [0,1,width,width+1]
  ]
  const iTetromino = [
    [1,width+1,width*2+1,width*3+1],
    [width,width+1,width+2,width+3],
    [1,width+1,width*2+1,width*3+1],
    [width,width+1,width+2,width+3]
  ]
  const theTetrominoes = [lTetromino, zTetromino, tTetromino, oTetromino, iTetromino]
  let currentPosition = 4
  let currentRotation = 0
  //randomly select a Tetromino and its first rotation
  let random = Math.floor(Math.random()*theTetrominoes.length)
  let current = theTetrominoes[random][currentRotation]
  let nextRandom = Math.floor(Math.random()*theTetrominoes.length)
  //draw the Tetromino
  function draw() {
    current.forEach(index => {
      squares[currentPosition + index].classList.add('tetromino')
      squares[currentPosition + index].style.backgroundColor = colors[random]
    })
  }
  //undraw the Tetromino
  function undraw() {
    current.forEach(index => {
      squares[currentPosition + index].classList.remove('tetromino')
      squares[currentPosition + index].style.backgroundColor = ''

    })
  }

  function control(e) {
    if (!timerId) {return}
    switch (e.keyCode) {
      case 37:
        moveLeft()
        break
      case 38:
        rotate()
        break
      case  39:
        moveRight()
        break
      case 40:
        moveInterval= 100
        break
      case 32:
        skipDown()
        break
      case 82:
        location.reload()
        break
    }
  }
  document.addEventListener('keydown', control);
  document.addEventListener('keyup', (e) => {if (e.keyCode === 40){moveInterval = 500}});
  let lastTime = 0;
  let moveInterval = 500
  //move down function
  function moveDown() {
    undraw()
    if (!isTaken()){
      currentPosition += width
    }
    draw()
    freeze()
  }

  function gameLoop(timestamp) {
    const deltaTime = timestamp - lastTime; // Calculate deltaTime
    if (deltaTime > moveInterval) {
      moveDown()
      lastTime = timestamp;
      if (gameOver()){
        return
      }
    }
    timerId = requestAnimationFrame(gameLoop);
  }

  function skipDown(){
    undraw()
    while (!isTaken()){
      currentPosition += width
    }
    draw()
    freeze()
  }
  //freeze function
  function freeze() {
    if(isTaken()) {
      current.forEach(index => squares[currentPosition + index].classList.add('taken'))
      //start a new tetromino falling
      random = nextRandom
      nextRandom = Math.floor(Math.random() * theTetrominoes.length)
      current = theTetrominoes[random][currentRotation]
      currentPosition = 4
      if (current.some(index => squares[currentPosition + index].classList.contains('taken'))){return true}
      draw()
      displayShape()
      addScore()
      return true
    }
    return false
  }

  function isTaken(){
    return current.some(index => squares[currentPosition + index + width].classList.contains('taken'))
  }

  //move the tetromino left, unless is at the edge or there is a blockage
  function moveLeft() {
    undraw()
    const isAtLeftEdge = current.some(index => (currentPosition + index) % width === 0)
    if(!isAtLeftEdge) currentPosition -=1
    if(current.some(index => squares[currentPosition + index].classList.contains('taken'))) {
      currentPosition +=1
    }
    draw()
  }
  //move the tetromino right, unless is at the edge or there is a blockage
  function moveRight() {
    undraw()
    const isAtRightEdge = current.some(index => (currentPosition + index) % width === width -1)
    if(!isAtRightEdge) currentPosition +=1
    if(current.some(index => squares[currentPosition + index].classList.contains('taken'))) {
      currentPosition -=1
    }
    draw()
  }

  function isAtRight() {
    return current.some(index=> (currentPosition + index + 1) % width === 0)  
  }

  function isAtLeft() {
    return current.some(index=> (currentPosition + index) % width === 0)
  }

  function checkRotatedPosition(P){
    P = P || currentPosition       //get current position.  Then, check if the piece is near the left side.
    if ((P+1) % width < 4) {         //add 1 because the position index can be 1 less than where the piece is (with how they are indexed).     
      if (isAtRight()){            //use actual position to check if it's flipped over to right side
        currentPosition += 1    //if so, add one to wrap it back around
        checkRotatedPosition(P) //check again.  Pass position from start, since long block might need to move more.
        }
    }
    else if (P % width > 5) {
      if (isAtLeft()){
        currentPosition -= 1
      checkRotatedPosition(P)
      }
    }
    if (current.some(index => squares[currentPosition + index].classList.contains('taken'))){
      currentRotation--
      //currentPosition -= width
      if(currentRotation === -1) { //if the current rotation gets to 4, make it go back to 0
        currentRotation = current.length
      }
      current = theTetrominoes[random][currentRotation]
    }
  }
  //rotate the tetromino
  function rotate() {
    undraw()
    currentRotation ++
    if(currentRotation === current.length) { //if the current rotation gets to 4, make it go back to 0
      currentRotation = 0
    }
    current = theTetrominoes[random][currentRotation]
    checkRotatedPosition()
    draw()
  }
  //show up-next tetromino in mini-grid display
  const displaySquares = document.querySelectorAll('.mini-grid div')
  const displayWidth = 4
  const displayIndex = 0
  //the Tetrominos without rotations
  const upNextTetrominoes = [
    [1, displayWidth+1, displayWidth*2+1, 2], //lTetromino
    [0, displayWidth, displayWidth+1, displayWidth*2+1], //zTetromino
    [1, displayWidth, displayWidth+1, displayWidth+2], //tTetromino
    [0, 1, displayWidth, displayWidth+1], //oTetromino
    [1, displayWidth+1, displayWidth*2+1, displayWidth*3+1] //iTetromino
  ]
  //display the shape in the mini-grid display
  function displayShape() {
    //remove any trace of a tetromino form the entire grid
    displaySquares.forEach(square => {
      square.classList.remove('tetromino')
      square.style.backgroundColor = ''
    })
    upNextTetrominoes[nextRandom].forEach( index => {
      displaySquares[displayIndex + index].classList.add('tetromino')
      displaySquares[displayIndex + index].style.backgroundColor = colors[nextRandom]
    })
  }
  //add functionality to the button
  startBtn.addEventListener('mousedown', () => {
    if (timerId) {
      cancelAnimationFrame(timerId)
      timerId = null
      stopStopwatch()
    } else {
      draw()
      timerId = requestAnimationFrame(gameLoop)
      displayShape()
      startStopwatch()
    }
  })

  restartbtn.addEventListener('mousedown', () => {
    location.reload();
  });

  function addScore() {
    for (let i = 0; i < 199; i +=width) {
      const row = [i, i+1, i+2, i+3, i+4, i+5, i+6, i+7, i+8, i+9]

      if(row.every(index => squares[index].classList.contains('taken'))) {
        score +=10
        scoreDisplay.innerHTML = score
        row.forEach(index => {
          squares[index].classList.remove('taken')
          squares[index].classList.remove('tetromino')
          squares[index].style.backgroundColor = ''
        })
        const squaresRemoved = squares.splice(i, width)
        squares = squaresRemoved.concat(squares)
        squares.forEach(cell => grid.appendChild(cell))
      }
    }
  }

  function gameOver() {
    if (current.some(index => squares[currentPosition + index].classList.contains('taken'))) {
      startBtn.disabled = true;
      scoreDisplay.innerHTML += ' End'
      cancelAnimationFrame(timerId)
      timerId = null
      stopStopwatch()
      //resetStopwatch()
      return true
    } else {
      return false
    }
  }
  var stopwatchElement = document.getElementById('stopwatch');
  var stopwatchTime = 0;
  var stopwatchInterval;
  function updateStopwatch() {
    var minutes = Math.floor(stopwatchTime / 60);
    var seconds = stopwatchTime % 60;
    stopwatchElement.textContent = `${minutes < 10 ? '0' : ''}${minutes}:${seconds < 10 ? '0' : ''}${seconds}`;
  }

  function startStopwatch() {
      stopwatchInterval = setInterval(function () {
          stopwatchTime++;
          updateStopwatch();
      }, 1000);
  }

  function stopStopwatch() {
      clearInterval(stopwatchInterval);
  }
  //saved for when not refreshing website for restart
  function resetStopwatch() {
      stopwatchTime = 0;
      updateStopwatch();
  }
})