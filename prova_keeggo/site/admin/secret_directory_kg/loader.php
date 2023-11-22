<?php 
    if (session_status() !== PHP_SESSION_ACTIVE) {
        session_start();
    }

    $sessionControl = false;

    if(isset($_SESSION['user'])){
        if ($_SESSION['user'] == "keeggo"){
            $sessionControl = true;
        }
    }

    if (!$sessionControl) header('Location: ../index.php');
?>

<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8" />
        <title>Desafio Keeggo</title>
    </head>
</html>

<?php

if(isset($_GET["pagina"])){


    if(isset($_SESSION['user'])){
        if ($_SESSION['user'] == "admin"){
            include $_GET["pagina"] . ".php";
        }
    } else{
        echo "Usuário não autenticado";
    }

}else{
    echo "Desculpe, não tem nada aqui...";
}

?>