let jadwal = JSON.parse(localStorage.getItem('jadwal')) || [];

const form = document.getElementById('jadwalForm');
const list = document.getElementById('jadwalList');

function simpanJadwal(event) {
  event.preventDefault();

  const data = {
    matkul: form.matkul.value,
    dosen: form.dosen.value,
    hari: form.hari.value,
    jam: form.jam.value,
    ruang: form.ruang.value
  };

  jadwal.push(data);
  localStorage.setItem('jadwal', JSON.stringify(jadwal));
  form.reset();
  tampilkanJadwal();
}

function tampilkanJadwal() {
  list.innerHTML = '';
  jadwal.forEach((item, index) => {
    list.innerHTML += `
      <tr>
        <td>${item.matkul}</td>
        <td>${item.dosen}</td>
        <td>${item.hari}</td>
        <td>${item.jam}</td>
        <td>${item.ruang}</td>
        <td>
          <button onclick="hapusJadwal(${index})">Hapus</button>
        </td>
      </tr>`;
  });
}

function hapusJadwal(index) {
  if (confirm('Yakin ingin menghapus?')) {
    jadwal.splice(index, 1);
    localStorage.setItem('jadwal', JSON.stringify(jadwal));
    tampilkanJadwal();
  }
}

form.addEventListener('submit', simpanJadwal);
tampilkanJadwal();
