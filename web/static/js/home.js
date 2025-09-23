const filterbar = document.getElementById('filterbar')
const filter = document.getElementById('filter')
const filteractiv= document.getElementById('filter filteractive')

filter.addEventListener('click', () => {
    container.classList.add("active");
});

filter.addEventListener('click', () => {
    container.classList.remove("active");
});

filterbar.addEventListener('click', () => {
    container.classList.remove("active");
});