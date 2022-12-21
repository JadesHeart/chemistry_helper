const urlBalance = 'http://localhost:8080/chemistry/balance';
const urlBase = 'http://localhost:8080/chemistry/base';
const urlTerm = 'http://localhost:8080/chemistry/term';
const urlPotential = 'http://localhost:8080/chemistry/potential';
const urlInstability = 'http://localhost:8080/chemistry/instability';
const tableBodyBalance = document.querySelector('#balance');
const tableBodyBase = document.querySelector('#base');
const tableBodyTerm = document.querySelector('#term');
const form = document.querySelector('.form');
const baseForm = document.querySelector('#baseForm')
const termForm = document.querySelector('#termForm')
const potentialForm = document.querySelector('#potentialForm')
const instabilityForm = document.querySelector('#instabilityForm')
potentialForm.addEventListener('submit', (e) => {
    e.preventDefault();

    const input = e.target.name;

    fetchData(urlPotential, input.value)
        .then((data) => {
            for (i=0;i<data.length;i++){
                document.querySelector('#potential'+i).innerHTML =addTableRowPotential(data[i]);
            }
            input.value = '';
        })
        .catch((err) => {
            document.getElementById('potential0').innerHTML = `
      <tr>
      <td colspan='4'>Неверное имя элемента ${err}</td>
      </tr>
      `;
        });
});
instabilityForm.addEventListener('submit', (e) => {
    e.preventDefault();

    const input = e.target.name;

    fetchData(urlInstability, input.value)
        .then((data) => {
            for (i=0;i<data.length;i++){
                document.querySelector('#instability'+i).innerHTML =addTableRowInstability(data[i]);
            }
            input.value = '';
        })
        .catch((err) => {
            document.getElementById('instability0').innerHTML = `
      <tr>
      <td colspan='4'>Неверное имя элемента ${err}</td>
      </tr>
      `;
        });
});
function addTableRowInstability(data) {
    return `
    <tr>
    <td>${data.ElName}</td>
    <td>${data.Ligand}</td>
    <td>${data.Complex}</td>
    <td>${data.LastParam}</td>
   </tr>`;
}
function addTableRowPotential(data) {
    return `
    <tr>
    <td>${data.Number}</td>
    <td>${data.Symbol}</td>
    <td>${data.ElName}</td>
    <td>${data.HalfReactions}</td>
    <td>${data.LastParam}</td>
   </tr>`;
}
termForm.addEventListener('submit', (e) => {
    e.preventDefault();

    const input = e.target.name;
    fetchData(urlTerm, input.value)
        .then((data) => {
            tableBodyTerm.innerHTML = addTableRowTerm(data);
            input.value = '';
        })
        .catch((err) => {
            tableBodyTerm.innerHTML = `
      <tr>
      <td colspan='4'>Неверное имя элемента ${err}</td>
      </tr>
      `;
        });
});
form.addEventListener('submit', (e) => {
    e.preventDefault();

    const input = e.target.name;

    fetchData(urlBalance, input.value)
        .then((data) => {
            tableBodyBalance.innerHTML = addTableRow(data);
            input.value = '';
        })
        .catch((err) => {
            tableBodyBalance.innerHTML = `
      <tr>
      <td colspan='4'>Неверное имя элемента ${err}</td>
      </tr>
      `;
        });
});
baseForm.addEventListener('submit', (e) => {
    e.preventDefault();

    const input = e.target.name;

    fetchData(urlBase, input.value)
        .then((data) => {
            tableBodyBase.innerHTML = addTableRowBase(data);
            input.value = '';
        })
        .catch((err) => {
            tableBodyBase.innerHTML = `
      <tr>
      <td colspan='4'>Неверное имя элемента ${err}</td>
      </tr>
      `;
        });
});
async function fetchData(url, param) {
    console.log(param)
    const res = await fetch(`${url}?name=${param}`);

    if (!res.ok) {
        const message = `An error has occured: ${res.status}`;
        throw Error(message);
    }

    return res.json();
}

function addTableRow(data) {
    return `
    <tr>
    <td>${data.ElName}</td>
    <td>${data.Formula}</td>
    <td>${data.FirstParam}</td>
    <td>${data.SecondParam}</td>
   </tr>
    `;
}
function addTableRowBase(data) {
    return `
    <tr>
    <td>${data.ElName}</td>
    <td>${data.ElectronicConfiguration}</td>
    <td>${data.StabilityOxidation}</td>
    <td>${data.MeltingPoint}</td>
    <td>${data.BoilingPoint}</td>
    <td>${data.ChemicalCompounds}</td>
   </tr>
    `;
}
function addTableRowTerm(data) {
    return `
    <tr>
    <td>${data.ElName}</td>
    <td>${data.Formula}</td>
    <td>${data.FirstParam}</td>
    <td>${data.SecondParam}</td>
    <td>${data.ThirdParam}</td>
    <td>${data.FourthParam}</td>
   </tr>
    `;
}