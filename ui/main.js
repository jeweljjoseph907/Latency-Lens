const output = document.getElementById('output');
const loadButton = document.getElementById('load');

loadButton.addEventListener('click', async () => {
  output.textContent = 'Loading…';
  try {
    const response = await fetch('http://localhost:9090/traces');
    const data = await response.json();
    output.textContent = JSON.stringify(data, null, 2);
  } catch (error) {
    output.textContent = `Request failed: ${error.message}`;
  }
});
