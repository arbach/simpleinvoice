$('#login-button').click(function(){
    $('#login-button').fadeOut("slow",function(){
        $("#container").fadeIn();
        TweenMax.from("#container", .4, { scale: 0, ease:Sine.easeInOut});
        TweenMax.to("#container", .4, { scale: 1, ease:Sine.easeInOut});
    });
});

$(".close-btn").click(function(){
    TweenMax.from("#container", .4, { scale: 1, ease:Sine.easeInOut});
    TweenMax.to("#container", .4, { left:"0px", scale: 0, ease:Sine.easeInOut});
    $("#container, #forgotten-container").fadeOut(800, function(){
        $("#login-button").fadeIn(800);
    });
});

/* Forgotten Password */
$('#forgotten').click(function(){
    $("#container").fadeOut(function(){
        $("#forgotten-container").fadeIn();
    });
});

function openForm2(){  

    var amount = document.querySelector("#f1 #amount");
    var desc = document.querySelector("#f1 #desc");

    invoice = {}
    invoice.amount = parseFloat(amount.value); 
    invoice.description = desc.value; 

    console.log(invoice);
    var url = window.location.protocol + "//" + window.location.host + "/invoice";

    $.ajax({
        url: url,
        type: 'post',
        dataType: 'json',
        contentType: 'application/json',
        success: function (data) {
            document.querySelector(".error").innerHTML = "";
            document.querySelector("#f2 #invoice_id").value = data.id;
            document.querySelector("#f2 #amount").value = data.amount;
            document.querySelector("#f2 #description").value = data.description;
            document.querySelector("#f2 #paid_amount").value = data.paidAmount;
            document.querySelector("#f2 #status").value = data.status;
            document.querySelector("#f2 #payment_address").value = data.paymentAddress;
        },
        data: JSON.stringify(invoice)
    });
    document.querySelector(".form1").style.display = "none"; 
    document.querySelector(".form2").style.display = "block"; 

}  
function openForm1(){

    document.querySelector(".form2").style.display = "none"; 
    document.querySelector(".form1").style.display = "block"; 

}
function checkStatus(){
  var invoiceId = document.querySelector("#f2 #invoice_id");
  let url = "/invoice?id=" + invoiceId.value;
  fetch(url).then((response) => {
    return response.json();
  }).then(data => {
    if(data.amount == undefined ){
      var err = "Could not find invoice by id: " + invoiceId.value;
      console.log(err);
      document.querySelector(".error").innerHTML = err;
      document.querySelector("#f2 #amount").value = "";
      document.querySelector("#f2 #description").value= "";
      document.querySelector("#f2 #paid_amount").value= "";
      document.querySelector("#f2 #status").value = "";
      document.querySelector("#f2 #payment_address").value = "";
    }else{
      document.querySelector(".error").innerHTML = "";
      document.querySelector("#f2 #amount").value = data.amount;
      document.querySelector("#f2 #description").value = data.description;
      document.querySelector("#f2 #paid_amount").value = data.paidAmount;
      document.querySelector("#f2 #status").value = data.status;
      document.querySelector("#f2 #payment_address").value = data.paymentAddress;
    }    
  }).catch(err =>{
    console.log("rejected", err)
  }) 
}

