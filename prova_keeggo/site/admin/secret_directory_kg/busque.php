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
<html lang="en">
    <html>
        <head>
            <meta charset="utf-8" />
            <title>Desafio Keeggo</title>
        </head>
    </html>
    <body width="100%" height="100%" style="display:flex; align-items: center; justify-content: center; flex-direction: column; margin-top: 5%;">
        <h1>Bem vindo ao sistema, <b>admin</b>!</h1>
        <img src="../../assets/images/keeggo.png" height="150px">
        <br>
        Faça a sua busca:
        <form action="busque.php" method="GET">
            <input type="text" name="buscar">
            <input type="submit" value="Buscar">
        </form>
    </body>
</html>

<style>
    input{
        padding: 5px;
        margin-top: 3px;
    }
</style>

<?php 

if(isset($_GET['buscar'])){

    if(isset($_SESSION['user'])){
        if ($_SESSION['user'] == "admin"){
            
            $mensagemId = rand(1,10);

            switch ($mensagemId) {
                case 1:
                    echo "</br>Não existe nada aqui...";
                    break;
                case 2:
                    echo "</br>Por favor não faça SQL injection...";
                    break;
                case 3:
                    echo "</br>Talvez...";
                    break;
                case 4:
                    echo "</br>Busque melhor, nosso sistema é bem protegido...";
                    break;
                case 5:
                    echo "</br>Mais afundo...";
                    break;
                case 6:
                    echo "</br>Hackes são proibídos aqui...";
                    break;
                case 7:
                    echo "</br>Nada...";
                    break;
                case 8:
                    echo "</br>Quase lá...";
                    break;
                case 9:
                    echo "</br>Vá embora...";
                    break;
                case 10:
                    echo "</br>Tente mais...";
                    break;
            }
        }
    } else{
        echo "Usuário não autenticado";
    }
    
}

?>

