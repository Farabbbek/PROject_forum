const container = document.getElementById('container')
const registrBtn = document.getElementById('register')
const loginBtn = document.getElementById('login')

registrBtn.addEventListener('click', () => {
    container.classList.add("active");
});

loginBtn.addEventListener('click', () => {
    container.classList.remove("active");
});