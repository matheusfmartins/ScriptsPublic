<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8" />
        <title>Desafio - Red Team SW</title>
        <link rel="icon" type="image/x-icon" href="../assets/images/red_team_sw.jpg">
    </head>
    <body>
        Faça a sua busca:
        <form action="busque.php" method="GET">
            <input type="text" name="buscar">
            <input type="submit" value="Buscar">
        </form>
    </body>
</html>

<?php 

if(isset($_GET['buscar'])){
    echo "Não existe nada aqui...";
}

?>

