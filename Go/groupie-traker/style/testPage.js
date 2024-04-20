var card1;
var card2;
var card3;
var card4;
var hiddenCards = [];
const card52 = document.getElementById("card52");
const cards = document.querySelectorAll(".card");
const redirect = document.getElementById("redirect");
const artistInfo = document.getElementById("artistInfo");
const tours = document.getElementById("tours");
const loading = document.getElementById("loading");
let delay = 0;
let cardNum = 1;
let contentNum = 1;
let i = 1;

//TODO event listener goes mad when div becomes display none
//TODO scrolling animation breaks on upscroll
//TODO add visual scrollbar to let people know where they are

//listens for page load
window.addEventListener("load", event => {

    //removes loading animation on page load
    loading.style.animation = "dissolve 200ms 1 normal forwards";
    loading.addEventListener("animationend", event => {
        loading.style.display = "none";
    });

    //listens for scrollwheel movement
    window.addEventListener("wheel", event => {
        //converts csrollwheel movement values to 1 and -1
        const delta = Math.sign(event.deltaY);
        switch (delta) {
            //downscroll
            case 1: {
                if (window.getComputedStyle(card52).display == "none") {
                    //hides cards
                    while (i <= 4) {
                        document.getElementById("card" + cardNum).style.animation = "";
                        document.getElementById("card" + cardNum).style.animation = "cardSlideOut 1s 1 normal";
                        document.getElementById("card" + cardNum).style.animationDelay = delay + "ms";
                        delay += 200;
                        cardNum++;
                        i++;
                    };
                    i = 1;
                    delay = 0;
                    cardNum--;
                    let lastCard = document.getElementById("card" + cardNum);
                    //listens for animation end on last card -> removes current card set -> displays next one
                    lastCard.addEventListener("animationend", () => {
                        cardNum = cardNum - 3;
                        while (i <= 4) {
                            document.getElementById("card" + contentNum).style.display = "none";
                            cardNum++;
                            i++;
                        };
                        i = 1;
                        while (i <= 4) {
                            document.getElementById("card" + contentNum).style.display = "flex";
                            cardNum++;
                            i++;
                        };
                        i = 1;
                        cardNum = cardNum - 4;
                    });
                };
                break;
            };
            //upscroll
            case -1: {
                if (window.getComputedStyle(document.getElementById("card1")) == "none") {
                    while (i <= 4) {
                        document.getElementById("card" + cardNum).style.animation = "";
                        document.getElementById("card" + cardNum).style.animation = "cardSlideOut 1s 1 normal";
                        document.getElementById("card" + cardNum).style.animationDelay = delay + "ms";
                        delay += 200;
                        cardNum++;
                        i++;
                    };
                    delay = 0;
                    cardNum--;
                    let lastCard = document.getElementById("card" + cardNum);
                    lastCard.addEventListener("animationend", () => {
                        cardNum = cardNum - 3;
                        while (i <= 4) {
                            document.getElementById("card" + contentNum).style.display = "none";
                            cardNum++;
                            i++;
                        };
                        i = 1;
                        cardNum = cardNum - 8;
                        while (i <= 4) {
                            document.getElementById("card" + contentNum).style.display = "flex";
                            cardNum++;
                            i++;
                        };
                        i = 1;
                    });
                };
                break;
            };
        };
    });

    //listens for click on "back" button -> displays cards again
    redirect.addEventListener("click", event => {
        removeInfo();
    });

    //listen for click on card -> displays artist info
    cards.forEach(card => {
        card.addEventListener("click", event => {
            //event.currentTarget is used, because event would just return child element where click was triggered
            let current = event.currentTarget;
            displayInfo(current);
        });
    });

    function displayInfo(current) {
        hideCards(current);
        appearInfo(current);
    };

    function hideCards(current) {
        let i = 1
        cards.forEach(card => {
            if (window.getComputedStyle(card).display == "block" && i != 4 && card != current) {
                card.style.animation = "";
                card.style.animation = "1s cardSlideOut 1 normal forwards";
                card.style.animationDelay = delay + "ms";
                delay += 200;
                i++;
            };
        });
        delay = 0;
    };

    function appearInfo(current) {
        cards.forEach(card => {
            if (window.getComputedStyle(card).display == "block" && card != current) {
                card.style.display = "none";
                hiddenCards.push(card)
            };
        });
        artistInfo.style.animation = "1s cardSlideInBot 0s 1 normal forwards";
        artistInfo.style.display = "flex";
        tours.style.opacity = "0";
        tours.style.animation = "1s cardSlideInBot 500ms 1 normal forwards";
        tours.style.display = "flex";
        redirect.style.opacity = "0";
        redirect.style.animation = "1s cardSlideIn 1s 1 normal forwards";
        redirect.style.display = "block";
    };

    function removeInfo() {
        hideInfo();
        redirect.addEventListener("animationend", showCards(hiddenCards));
    };

    function hideInfo() {
        artistInfo.style.animation = "1s cardSlideOut 0ms 1 normal forwards";
        tours.style.opacity = "1";
        tours.style.animation = "1s cardSlideOut 500ms 1 normal forwards";
        redirect.style.opacity = "1";
        redirect.style.animation = "1s cardSlideOutUp 500ms 1 normal forwards";
    };

    function showCards() {
        artistInfo.style.display = "none";
        tours.style.display = "none";
        cards.forEach(card => {
            if (card == hiddenCards[0] || card == hiddenCards[1] || card == hiddenCards[2]) {
                card.style.animation = "";
                delay += 200;
                card.style.animation = "1s cardSlideIn 1 normal forwards";
                card.style.animationDelay = delay + "ms";
                card.style.display = "block";
            };
        });
    };

});