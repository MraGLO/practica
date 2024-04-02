export const getData = async (url) => {
    const response = await fetch(url);

    if (!response.ok) {
        throw new Error(`Ошибка по адресу ${url}, статус ошибки ${response}`);
    }

    return await response.json();
};

export const sendData = async (url, data) => {
    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)

        });

        if (!response.ok) {
            document.querySelector('.message__container').innerHTML += 
            `
            <div class="message__box message__box-error">
                <strong class="message__header">Error:</strong>
                <div class="message__text">${response.status}:${response.statusText}</div>
            </div>
            `
            throw new Error(`Ошибка по адресу ${url}, статус ошибки ${response.status}: ${response.statusText}`);

        }
        else{
            document.querySelector('.message__container').innerHTML += 
            `
            <div class="message__box message__box-success">
                <strong class="message__header">Success</strong>
                <div class="message__text">${response.status}:${response.statusText}</div>
            </div>
            `
        }


        return await response;
    } catch (error) {
        console.error('Ошибка при отправке данных:', error);
        throw error;
    }
};