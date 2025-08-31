function onSelect(result) {
    $("#email").val(result.email);
    $("#email").prop('disabled', true);;
}


function getFrmDat() {
    email = $("#email").val();
    return {
      "email": email};
}


function onNew() {
    $("#email").val("");
    $("#email").prop('disabled', false);;
}


function submitValidation() {
    return true;
}

$(document).ready(function() {
    $("#divInputBtn").css("margin-left", "calc(100% - 80px)");
    $("#divInputBtn").css("width", "80px");
    $("#btnCancel").css("display", "none");
    InitTbl(0, onSelect);
    NewProcess(onNew);
    SubmitProcess(submitValidation, getFrmDat);
});
