<?php
    session_start();

    include ("dp.php");

    if($_SERVER['REQUEST_METHOD'] == 'POST')
    {
        $fname = $_POST['fname'];
        $lname = $_POST['lname'];
        $gender = $_POST['gender'];
        $gmail = $_POST['mail'];
        $password = $_POST['pass'];

        if(!empty($gmail) && !empty($password) && !is_numeric($gmail))
        {        
            
            $query = "insert into users (first_name, lname, gender, email, pass) values ('$fname', '$lname', '$gender', '$gmail', '$password')";

            mysqli_query($con, $query);

            echo "<script type='text/javascript'> alert('Successfully Register')</script>";
        }
        else
        {
            echo "<script type='text/javascript'> alert('Please enter some valid information!')</script>";
        }

     }
?>

<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width-device-width, initial-scale=1">
    <title>Form Login and Register</title>
    <link rel="stylesheet" type="text/css" href="style.css">

</head>
<body>
    <div class="signup">
        <h1>Sign Up</h1>
        <h4>It's free and only take a minute</h4>
        <form method="POST">
            <label>First Name</label>
            <input type="text" name="fname" required>
            <label>Last Name</label>
            <input type="text" name="lname" required>
            <label>Gender</label>
            <input type="text" name="gender" required>
            <label>Email</label>
            <input type="email" name="mail" required>
            <label>Password</label>
            <input type="password" name="pass" required>
            <input type="submit" name="" value="Submit">
        </form>
        <p>By clicking the Sign Up button, you agree to our<br>
        <a href="">Terms and Condition</a> and <a href="#">Policy Privacy</a>
        </p>
        <p>Already have an account? <a href="login.php">Login</a></p>
    </div>
</body>
</html>
