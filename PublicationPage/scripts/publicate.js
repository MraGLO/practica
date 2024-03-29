import { sendData } from './server.js';

//сбор выбранных элементов списка
function collectSelectedOptions(selectName) {
    const selectedItems = document.querySelectorAll(`.selected__list[name="${selectName}"] .selected__list-item`);
    const selectedValues = Array.from(selectedItems).map(item => item.value);
    return selectedValues;
}

// отправка на сервер
const publicate = (url) => {
    const publicationForm = document.querySelector('.publication-form');

    publicationForm.addEventListener('submit', e => {
        e.preventDefault();

       // Собираем данные из формы в объект
       const formDataObj = {};

       // Сбор данных из формы
       const formData = new FormData(publicationForm);
       for (let [key, value] of formData.entries()) {
           formDataObj[key] = value;
       }

        var categories = collectSelectedOptions('category');
        var tags = collectSelectedOptions('tag');
        categories = categories.map(category => Number(category));
        tags = tags.map(tag => Number(tag));

        // Добавляем данные в объект
        formDataObj.category = categories;
        formDataObj.tag = tags;

        // Отправляем данные на сервер
        sendData(url, formDataObj)
            .then(() => {
                publicationForm.reset();
                clearSelectedOptions();
            })
            .catch((err) => {
                console.log(err);
            });
    });
    publicationForm.addEventListener('reset', e => {
        publicationForm.reset();
        clearSelectedOptions();
    });

    function clearSelectedOptions() {
        document.querySelectorAll('.selected__list').forEach(selectedList => {
            selectedList.innerHTML = ''; 
        });
    }
}

publicate('/newNews/add');