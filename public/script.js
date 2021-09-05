let out, xhr;
xhr = new XMLHttpRequest;

window.onload = function() {

}

function on_page(name) {
	xhr.open("GET", "public/pages/"+name, true);
	xhr.send();
	xhr.onreadystatechange = function() {
		if(xhr.readyState != 4) return;
		if (xhr.status == 404) {
			return document.querySelector("body").innerHTML = `
				<h1>Ошибка 404</h1>
				<h3>Такой страницы не существует</h3>
				<a href="/">На главную</a>
			`;
		}
		document.querySelector("#app").innerHTML = xhr.responseText;
	}
}

function get(callback, data, url) {
	xhr.open("GET", url, true);
	xhr.send(data);
	xhr.onreadystatechange = function() {
		if(xhr.readyState != 4) return;
		callback(xhr.responseText);
	}
}

function post(callback, data, url) {
	xhr.open("GET", url, true);
	xhr.setRequestHeader('Content-Type', 'application/json');
	xhr.send(data);
	xhr.onreadystatechange = function() {
		if(xhr.readyState != 4) return;
		callback(xhr.responseText);
	}
}

function get_teachers() {
	get(data => {
		data = JSON.parse(data); out = "";
		data.forEach(teacher => {
			out += `
				<tr>
					<td>${teacher.ID}</td>
					<td>${teacher.Surname}</td>
					<td>${teacher.Name}</td>
					<td>${teacher.Patronymic}</td>
					<td>${teacher.Post}</td>
					<td>${teacher.Education}</td>
					<td>${teacher.Qualification}</td>
				</tr>
			`;
		});
		document.querySelector("tbody").innerHTML = out;
	}, null, "api/teachers");
}

function get_students() {
	get(data => {
		data = JSON.parse(data); out = "";
		data.forEach(student => {
			out += `
				<tr>
					<td>${student.ID}</td>
					<td>${student.Surname}</td>
					<td>${student.Name}</td>
					<td>${student.Patronymic}</td>
					<td>${student.DateBirth}</td>
					<td>${student.ReceiptDate}</td>
					<td>${student.ExperationDate}</td>
				</tr>
			`;
		});
		document.querySelector("tbody").innerHTML = out;
	}, null, "api/students");
}