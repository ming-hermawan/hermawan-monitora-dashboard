function onSelect(result) {
    $("#username").val(result.username);
    $("#username").prop('disabled', true);
    $("#username0").val(result.username);
    $("#name").val(result.name);
    $("#name0").val(result.name);
    $("#email").val(result.email);
    $("#email0").val(result.email);
    $("#isAdmin").val(result.isAdmin);
    $("#isAdmin0").val(result.isAdmin);
    $("#isAdmin").prop("checked", (result.isAdmin === 1));
    initSelect("group", result.usergroup, result.usergroupList, "");
    $("#group0").val(result.usergroup);
}


function onNew(e) {
    $.ajax({
      url: "/ddl/admin/usermanagement/user-group",
      type: "POST",
      dataType: "json",
      success: function (result) {
        initSelect("group", null, result, "");
      }
    });
    $("#username").prop('disabled', false);
    $("#username").val("");
    $("#username0").val("");
    $("#name").val("");
    $("#name0").val("");
    $("#isAdmin").prop("checked", false);
    $("#isAdmin").val(0);
    $("#isAdmin0").val(0);
    $("#email").val("");
    $("#email0").val("");
    $("#divInput").css("display", "block");
}


function onCancel(e) {
    $("#usergroup").val($("#usergroup0").val());
    $("#password").val($("#password0").val());
    $("#name").val($("#name0").val());
    $("#email").val($("#email0").val());
    $("#isAdmin").val($("#isAdmin0").val());
    $("#isAdmin").prop("checked", ($("#isAdmin0").val() === 'true'));
    $("#group").val($("#group0").val());
}

function submitValidation() {
    var password = $("#password").val();
    var passwordconfirm = $("#password").val();
    if((password === "") || (password === passwordconfirm))
        return true;
    else
        return false;
}

function getFrmDat() {
    let username = $("#username").val();
    let name = $("#name").val();
    let is_admin = $("#isAdmin").prop("checked");
    if(is_admin) {
        val_is_admin = 1;
    } else {
        val_is_admin = 0;
    }
    let password = $("#password").val();
    let email = $("#email").val();
    let group = $("#group").find(":selected").text();
    var out = {
      "username": username,
      "name": name,
      "isAdmin": val_is_admin,
      "email": email,
      "group": group};
    if(password !== "") {
        out["password"] = password;
    }
    return out;
}


$(document).ready(function(){
    InitTbl(0, onSelect);
    NewProcess(onNew);
    CancelProcess(onCancel);
    SubmitProcess(submitValidation, getFrmDat);

    //$("#frmInp").submit(onSubmit);

    $("#password").change(function(){
        if($(this).val().trim() !== "") {
            $("#passwordconfirm").prop('disabled', false);
        }
    });
});
