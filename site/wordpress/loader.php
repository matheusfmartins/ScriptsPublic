<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8" />
        <title>Desafio - Red Team SW</title>
        <link rel="icon" type="image/x-icon" href="../assets/images/red_team_sw.jpg">
    </head>
</html>

<?php

if(isset($_GET["pagina"])){
    include $_GET["pagina"] . ".php";
}else{
    echo "Desculpe, não tem nada aqui...";
}

?>