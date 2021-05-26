const { event } = require("jquery");

$('#transform-Json').on('submit', transformJson);

function transformJson(evento){
    evento.preventDefault();
    console.log("Dentro da função transformation");

    if($('#request').val().length != 0){
        alert("Requisição não encaminhada")
        return
    }


    if (transformation != ""){
        alert("Transformação não encaminhada")
        return
    }


    $.ajax({
        url:"transformation",
        method:"POST",
        data:{
            transformation: $('#transformation').val(),
            request : $('#request').val()
        }
    })
}