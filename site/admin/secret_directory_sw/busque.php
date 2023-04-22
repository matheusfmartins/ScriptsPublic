<?php 
    if (session_status() !== PHP_SESSION_ACTIVE) {
        session_start();
    }

    $sessionControl = false;

    if(isset($_SESSION['user'])){
        if ($_SESSION['user'] == "admin"){
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
            <title>Desafio - Red Team SW</title>
            <link rel="icon" type="image/x-icon" href="../../assets/images/red_team_sw.jpg">
        </head>
    </html>
    <body width="100%" height="100%" style="display:flex; align-items: center; justify-content: center; flex-direction: column; margin-top: 5%;">
        <h1>Buscar</h1>
        <img src="../../assets/images/red_team_sw.jpg">
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
                    echo "Não existe nada aqui...";
                    break;
                case 2:
                    echo "Por favor não faça SQL injection...";
                    break;
                case 3:
                    echo "Talvez...";
                    break;
                case 4:
                    echo "Busque melhor, nosso sistema é bem protegido...";
                    break;
                case 5:
                    echo "Mais afundo...";
                    break;
                case 6:
                    echo "Hackes são proibídos...";
                    break;
                case 7:
                    echo "Nada...";
                    break;
                case 8:
                    echo "Quase lá...";
                    break;
                case 9:
                    echo "Vá embora...";
                    break;
                case 10:
                    echo "Tente mais...";
                    break;
            }
        }
    } else{
        echo "Usuário não autenticado";
    }
    
}

?>

