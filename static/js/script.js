function getFiles() {
    document.getElementById("allFiles").innerHTML = "";
    document.getElementById("selectedFiles").innerHTML = "";
    document.getElementById("uploadBtn").setAttribute("disabled", "disabled")
    $.ajax({
        method: "GET",
        url: "http://localhost:8080/files",
        success: function (res) {
            for (var i = 0; i < res.length; i += 1) {
                var divFile = document.createElement('div');
                divFile.setAttribute("class", "file");
                var divForLink = document.createElement("div");
                divForLink.setAttribute("class", "fileLink");
                var a = document.createElement('a');
                a.innerHTML = res[i].FileName
                divFile.appendChild(divForLink)
                divForLink.appendChild(a);
                a.setAttribute("id", res[i].ID)
                var div1 = document.createElement('div');
                div1.setAttribute("class", "generateLink");
                const button = document.createElement("button");
                button.textContent = "Сгенерировать одноразовую ссылку для скачивания"
                button.setAttribute("class", "buttonLink")
                const fileId = res[i].ID
                button.onclick = function () {
                    $.ajax({
                        method: "GET",
                        url: "http://localhost:8080/createToken?id=" + fileId,
                        success: function (res) {
                            alert("Одноразовая ссылка на скачивание: " + "http://localhost:8080/downloadByToken?token=" + res.token)
                        },
                        error: function (res) {
                            alert("Не удалось получить ссылку на скачивание.")
                        }
                    })
                }
                divFile.appendChild(div1)
                div1.appendChild(button)
                req = "http://localhost:8080/download?id=" + res[i].ID
                a.setAttribute("href", req)
                document.getElementById("allFiles").appendChild(divFile);
            }
        }

    })
}

function uploadFiles() {
    fileContainer.forEach((file, id) => {
        const p = document.createElement("p");
        p.innerText = "загружается..."
        p.setAttribute("id", "fileUploadStatus" + id)
        document.getElementById("selectedFile" + id).appendChild(p)
        var formData = new FormData();
        formData.append("myFile", file)
        $.ajax({
            method: "POST",
            url: "http://localhost:8080/upload",
            data: formData,
            async: true,
            processData: false,
            contentType: false,
            success: function (res) {
                document.getElementById("fileUploadStatus" + id).innerText = "Файл успешно загружен!"
            },
            error: function (res) {
                document.getElementById("fileUploadStatus" + id).innerText = "Не удалось загрузить файл!"
            }
        })
    })
    fileContainer = [];
}

var fileContainer = [];

function selectFile() {
    document.getElementById("allFiles").innerHTML = "";
    var input = document.createElement('input');
    input.type = 'file';
    input.setAttribute("multiple", "multiple")

    input.onchange = e => {
        for (let i = 0; i < e.target.files.length; i++) {
            fileContainer.push(e.target.files[i]);
        }
        showFileContainer()
    }
    input.click();
}

function showFileContainer() {
    const selectedFiles = document.getElementById("selectedFiles")
    selectedFiles.innerHTML = ""
    fileContainer.forEach(((file, id) => {
        var li = document.createElement('li');
        li.setAttribute("id", "selectedFile" + id)
        var p = document.createElement('p');
        p.innerHTML = file.name
        li.appendChild(p);
        selectedFiles.appendChild(li);
    }))
    if (fileContainer.length > 0) {
        document.getElementById("uploadBtn").removeAttribute("disabled")
    } else {
        document.getElementById("uploadBtn").setAttribute("disabled", "disabled")
    }

}