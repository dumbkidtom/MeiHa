<html>
<title>lunch!</title>
<body>

<style>

  html {
    font-family: Sans-Serif;
  }

  .card {
    border-radius: 5px;
    box-shadow: 0 4px 8px 0 rgba(0,0,0,0.2);
    transition: 0.3s;
    width: 300px;
    background-color: #33BFFF;

    margin-left: auto;
    margin-right: auto;
    margin-top: 5%;
  }

  .card:hover {
    box-shadow: 0 8px 16px 0 rgba(0,0,0,0.2);
  }

  .container {
    padding: 2px 16px;
    background-color: white;
  }

  input[type="submit"] {
    background: none;
    box-shadow: none;
    border-radius: 0px;
    border: 0;
  }
</style>


  <div class="card">
    <table border=0 cellspacing=0 cellpadding=10px>
    <tr><td><font size=15px><b>Lunch</b></font><br>
            </td></tr>
    </table>


  <div class="container">
  <P>
  <form method=post action=/search>
    <table cellspacing=10 border=0>
     <tr><td>ZIP:</td><td><input name=zip value="77042"></td></tr>
     <tr><td>RADIUS:</td><td><select name=radius>
         <option value=1>1</option>
         <option value=5>5</option>
         <option value=10>10</option>
         <option value=15>15</option></select> miles</td></tr>
    </table>
    <p align=right>
    <input type=submit value="GO"></right>
    </p>
  </form>
  </div>
  </div>
</body>
</html>
