import { sendData } from './server.js';

//сбор выбранных элементов списка
function collectSelectedOptions(selectName) {
    // Используем класс для поиска выбранных элементов, а не имя элемента ul
    const selectedItems = document.querySelectorAll(`.selected__list[name="${selectName}"] .selected__list-item`);
    const selectedValues = Array.from(selectedItems).map(item => item.value);
    return selectedValues;
}

// отправка на сервер
const publicate = (url) => {
    const publicationForm = document.querySelector('.publication-form');

    publicationForm.addEventListener('submit', e => {
        e.preventDefault();

        const formData = new FormData(publicationForm);

        const categories = collectSelectedOptions('categories');
        const tags = collectSelectedOptions('tags');
        formData.set('categories', JSON.stringify(categories));
        formData.set('tags', JSON.stringify(tags));

        sendData(url, formData)
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

publicate('Post/news');    