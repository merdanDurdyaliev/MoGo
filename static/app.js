
let openNavSpan = document.getElementById('span');
let mySideNav = document.getElementById('mySidenav');
let closeButton = document.getElementById('closebtn');
let openAuth = document.getElementById('openAuth');
let closeAuth = document.getElementById('closeAuth');
let autho = document.getElementById('autho');
let registrate = document.getElementById('registrate');
let creatAcc = document.getElementById('creatAcc');
let signIn = document.getElementById('signIn');
let signInto = document.getElementById('signInto');
let user = document.getElementById('User');

registrate.onclick = function openCreatAcc() {
    creatAcc.style.display = "grid";
    signIn.style.display = "none";
}

signInto.onclick = function openSignIn() {
    creatAcc.style.display = "none";
    signIn.style.display = "grid";
}

span.onclick = function openNav() {
    mySideNav.style.width = "100%";
    span.style.display = "none";
}

closeButton.onclick = function closeNav() {
    mySideNav.style.width = "0";
    span.style.display = "block";
}

openAuth.onclick = function openAuth() {
    autho.style.display = "block";
    mySideNav.style.width = "0";
}

user.onclick = function openAutho() {
    autho.style.display = "block";
    mySideNav.style.width = "0";
}

closeAuth.onclick = function closeAuth() {
    autho.style.display = "none";
}



let accordionItems = document.getElementsByClassName('accordion__item');
let previewButtons = document.getElementsByClassName('reviews__btn');
let reviewsItems = document.getElementsByClassName ('reviews__item');
let reviewsIts = document.getElementsByClassName ('reviews__it');

for (let i = 0; i < accordionItems.length; i++) {
    accordionItems[i].addEventListener('click', e => {
        if (e.target.classList.contains('accordion__header-button')) {
            accordionItems[i].classList.toggle('active');
        }
    })
}


for (let i = 0; i < previewButtons.length; i++) {
    let currentBtn = previewButtons[i];

    currentBtn.addEventListener('click', e => {

        for (let j = 0; j < reviewsItems.length; j++) {
            let currentItem = reviewsItems[j]

            if (e.target.classList.contains('reviews__btn--prev')) {
                if (currentItem.classList.contains('active') && j != 0) {
                    currentItem.classList.remove('active');
                    reviewsItems[j - 1].classList.add('active')
                    return
                }
            }

            if (e.target.classList.contains('reviews__btn--next')) {
                if (currentItem.classList.contains('active') && j != reviewsItems.length - 1) {
                    currentItem.classList.remove('active');
                    reviewsItems[j + 1].classList.add('active')
                    return
                }
            }
        }
    })
}



for (let i = 0; i < previewButtons.length; i++) {
    let currentBtn = previewButtons[i];

    currentBtn.addEventListener('click', e => {
        for (let j = 0; j < reviewsIts.length; j++){
            let currentItem = reviewsIts[j];

            if(e.target.classList.contains('reviews__btn--prev')) {
                if (currentItem.classList.contains('active') && j != 0) {
                    currentItem.classList.remove('active');
                    reviewsIts[j - 1].classList.add('active')
                    return
                }
            }


            if (e.target.classList.contains('reviews__btn--next')) {
                if (currentItem.classList.contains('active') && j != reviewsIts.length - 1){
                    currentItem.classList.remove('active');
                    reviewsIts[j + 1].classList.add('active')
                    return
                }
            }
        }
    })
}



