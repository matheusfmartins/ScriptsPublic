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

<?php
$servername = "localhost";
$database = "desafio_redteam_sw";
$username = "redteamsw";
$password = "y0um4yh4ckm3";
// Create connection
$conn = mysqli_connect($servername, $username, $password, $database);
// Check connection
if (!$conn) {
    die("Connection failed: " . mysqli_connect_error());
}
echo "Connected successfully";
mysqli_close($conn);
?>
