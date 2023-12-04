function enviarRequisicao() {
    // Obtenha os dados do formulário
    var formulario = document.getElementById("dados_cadastro");
    var formData = new FormData(formulario);

    // Converta os dados para um objeto JavaScript
    var jsonData = {};
    formData.forEach(function(value, key){
        jsonData[key] = value;
    });

    // Configure a URL do endpoint
    var url = "http://localhost:3000/api/user";

    // Opções da requisição
    var options = {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(jsonData)
    };

    // Faça a requisição usando a API Fetch
    fetch(url, options)
        .then(response => response.json())
        .then(data => {
            // Faça algo com a resposta do servidor, se necessário
            console.log(data);
        })
        .catch(error => {
            // Trate erros, se houver
            console.error('Erro ao enviar requisição:', error);
        });
}
