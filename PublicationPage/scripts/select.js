import { getData } from './server.js';  

async function parseData(url, selectName) {
    const data = await getData(url);
    const selectElement = document.querySelector(`ul[name="${selectName}"]`);

    data.forEach(item => {
        selectElement.innerHTML += `
        <li class="dropdown__list-item" value="${item.id}">${item[`${selectName}Name`]}</li>
        `;
    });
    sortListItems(selectElement);
}

// Функция для сортировки элементов списка по значению атрибута value
function sortListItems(listElement) {
    Array.from(listElement.children)
        .sort((a, b) => a.getAttribute('value') - b.getAttribute('value'))
        .forEach(item => listElement.appendChild(item));
}

// Функция для добавления обработчика событий к кнопкам внутри выбранного списка
function addButtonClickHandler(button, clickedItem, listName) {
    button.addEventListener('click', function() {
        // Возвращаем элемент в список
        const newListItem = document.createElement('li');
        newListItem.classList.add('dropdown__list-item');
        newListItem.setAttribute('value', clickedItem.getAttribute('value'));
        newListItem.innerText = clickedItem.innerText;
        const listElement = document.querySelector(`ul[name="${listName}"]`);
        listElement.appendChild(newListItem);

        // Сортируем список после добавления элемента
        sortListItems(listElement);

        // Удаляем кнопку
        this.remove();
    });
}



// Выбор списка
document.querySelectorAll('.dropdown').forEach(function (dropdownWrapper) {
    dropdownWrapper.querySelector('.dropdown__input').addEventListener('click', function () {
        dropdownWrapper.querySelector('.dropdown__list').classList.add('dropdown__list--visible');
    });

    dropdownWrapper.querySelector('.dropdown__input').addEventListener('input', function () {
        console.log(this.value)
        if (this.value.trim() === '') {

        }
        else {

        }
    });

    // Переключение видимости списка
    document.addEventListener('click', function (e) {
        const dropdownInput = dropdownWrapper.querySelector('.dropdown__input');
        const dropdownList = dropdownWrapper.querySelector('.dropdown__list');

        if (!dropdownInput.contains(e.target) && 
        !e.target.closest('.dropdown__list-item') && 
        dropdownWrapper.querySelector('.dropdown__list').classList.contains('dropdown__list--visible')) {
            dropdownList.classList.toggle('dropdown__list--visible');
        }
    });

    // Добавление кнопок в зависимости от списка
    dropdownWrapper.querySelector('.dropdown__list').addEventListener('click', function (e) {
        if (e.target.classList.contains('dropdown__list-item')) {
            const clickedItem = e.target;
            const listName = this.getAttribute('name');
            const selectedListSelector = `.selected__list[name="${listName}"]`;
            const selectedList = document.querySelector(selectedListSelector);

            // Создаем кнопку и добавляем ее в выбранный список
            const newButton = document.createElement('button');
            newButton.classList.add('selected__list-item');
            newButton.setAttribute('value', clickedItem.getAttribute('value'));
            newButton.innerText = clickedItem.innerText;
            selectedList.appendChild(newButton);

            // Удаляем выбранный элемент из списка
            clickedItem.remove();

            // Добавляем обработчик событий для кнопки
            addButtonClickHandler(newButton, clickedItem, listName);
        }
    });
});

parseData('/category/all', 'category');
parseData('/tag/all', 'tag');