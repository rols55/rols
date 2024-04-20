//content block number
let contentNum = 1;
//card nr
let cardNum = 1;
//inits aniation delay value
let delay = 0;
//lastCard element is used to listen for the animations on last card
//and switching content when animation plays out
let lastCard = document.getElementById("card" + cardNum);

window.addEventListener("DOMContentLoaded",
    //listens to mousewheel scrolls and makes them appear as 1 and -1
    window.addEventListener("wheel", event => {
        //inits last card in the content box for use in checking if the the card animation on age has been finished, 
        //because last card always gets more delay
        let lastCard = document.getElementById("card" + cardNum);
        console.log("card" + cardNum)
        console.log(lastCard);
        //simplifies scrollwheel values to 1 and -1
        const delta = Math.sign(event.deltaY);
        console.info(delta);
        switch (delta) {
            //positive value means, that mouse was scrolled downwards
            //scrolldown applies slide out animation to cards and after that it swithces to another set of cards
            case 1:
                console.log("downscroll");
                //because we have 13 sets of cards we need to check, that we don't go over them
                if (contentNum <= 12) {
                    console.log("in downscroll");
                    //applies slide out animation to cards, since cards default animation of sliding in already played
                    for (let i = 1; i <= 4; i++) {
                        console.log("in while down");
                        document.getElementById("card" + cardNum).style.animation = "";
                        document.getElementById("card" + cardNum).style.animation = "cardSlideOut 1s 1 normal";
                        document.getElementById("card" + cardNum).style.animationDelay = delay + "ms";
                        delay = + 200;
                        cardNum++;
                        console.log("card at whiles end" + cardNum);
                    }
                    //decrements card number because after loop last cards value becomes first cards from next card set
                    cardNum--;
                    console.log(cardNum + " after setting animations");
                    //gets last cards element into a variable
                    lastCard = document.getElementById("card" + cardNum);
                    console.log(lastCard);
                    //listens to the the animation on last card
                    //if the animation played out it means that we should switch to next set of cards 
                    lastCard.addEventListener("animationend", () => {
                        let b = window.getComputedStyle(lastCard);
                        console.log("card in animaiton " + cardNum);
                        console.log("animation is " + lastCard.getAttribute("opacity"));
                        console.log("listen to " + lastCard);
                        console.log("content to disappear" + contentNum);
                        //makes current set of cards invisible
                        document.getElementById("content" + contentNum).style.display = "none";
                        //switches to next set of cards
                        contentNum++;
                        //displays relevant set of cards
                        document.getElementById("content" + contentNum).style.display = "flex";
                        console.log("content to appear" + contentNum);
                        cardNum++;
                        console.log("after animation card nr is " + cardNum + " and card is " + lastCard);
                        console.log(lastCard);
                        return;
                    });
                    //resets animation delay property
                    console.log(cardNum + " after displaying new content");
                    delay = 0;
                    console.log("content displaying" + contentNum);
                    console.log("afer displaying new content");
                    break;
                } else if (contentNum == 13) {
                    break;
                }
            //negative value means that mouse was scrolled upwards
            case -1:
                console.log("upscroll");
                //our sets of cards are staring form 1
                if (contentNum >= 1) {
                    console.log("in upscroll");
                    //applies slide out animation for current cards, since slide plays on its own, because it's the default card animation
                    for (let i = 1; i <= 4; i++) {
                        console.log("in while up");
                        document.getElementById("card" + cardNum).style.animation = "";
                        document.getElementById("card" + cardNum).style.animation = "cardSlideOut 1s 1 normal";
                        document.getElementById("card" + cardNum).style.animationDelay = delay + "ms";
                        delay = + 200;
                        cardNum++;
                        console.log("card at whiles end" + cardNum);
                    }
                    //same behaviour as with scrolldown
                    cardNum--;
                    console.log(cardNum + " after setting animations");
                    lastCard = document.getElementById("card" + cardNum);
                    console.log(lastCard);
                    lastCard.addEventListener("animationend", () => {
                        let a = window.getComputedStyle(lastCard);
                        console.log("card in animaiton " + cardNum);
                        console.log("animation is " + lastCard.getAttribute("opacity"));
                        console.log("listen to " + lastCard);
                        console.log("content to disappear" + contentNum);
                        document.getElementById("content" + contentNum).style.display = "none";
                        //decrements sets of cards because we are going up the page to previous set
                        contentNum--;
                        document.getElementById("content" + contentNum).style.display = "flex";
                        console.log("content to appear" + contentNum);
                        cardNum - 7;
                        console.log("after animation card nr is " + cardNum + " and card is " + lastCard);
                        console.log(lastCard);
                        return; 
                    });
                    console.log(cardNum + " after displaying new content");
                    delay = 0;
                    console.log("content displaying" + contentNum);
                    console.log("afer displaying new content");
                    break;
                } else if (contentNum == 1) {
                    break;
                }
        }
    }
    )
);
