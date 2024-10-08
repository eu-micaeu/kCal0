import { openOverlay } from './functions/openOverlay.mjs';
import { closeOverlay } from './functions/closeOverlay.mjs';
import { loadMenu } from './functions/loadMenu.mjs';
import { btCreate } from './functions/btCreate.mjs';
import { btLogout } from './functions/btLogout.mjs';

document.addEventListener("DOMContentLoaded", function () {

    btCreate();

    btLogout();

    loadMenu();

    var spin = document.getElementById('spin');

    var main = document.getElementsByTagName('main')[0];

    setTimeout(function () {

        main.style.display = 'flex'; 

        spin.style.display = 'none';
          
    }, 2000); 

});

// Lógica de adicionar refeição
document.getElementById('btAddMeal').addEventListener('click', function () {

    openOverlay('popUpCreateMeal');

});


document.getElementById('divOverlay').addEventListener('click', function () {

    closeOverlay('popUpCreateMeal');

    closeOverlay('popUpCreateFood');

    closeOverlay('popUpCreateMenu');

});