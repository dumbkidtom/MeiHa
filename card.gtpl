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

  input[type="button"] {
    background: none;
    box-shadow: none;
    border-radius: 0px;
    border: 0;
  }
</style>

  <div class="card">
    <table border=0 cellspacing=0 cellpadding=10px>
    <tr><td><b>{{.Name}}</b><br>
            </td></tr>
    </table>

    <div class="container">
      <form method=GET>
      <P>
        <table cellspacing=10 border=0>
         <tr><td><a href="{{.URL}}"><img src="{{.IMGURL}}" width="250"></a></td></tr>
        </table>
        <p align=right>
        <input type=button value="Next" onclick="window.location.reload()"></right>
        </p>
      </form>
    </div>

  </div>

</body>
</html>