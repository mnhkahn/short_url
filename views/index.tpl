<!DOCTYPE html>

<html>
  	<head>
    	<title>Cyeam Shorten Url</title>
    	<link rel="shortcut icon" href="/static/c32.ico">

    	<link href="//maxcdn.bootstrapcdn.com/bootstrap/3.2.0/css/bootstrap.min.css" rel="stylesheet">

	</head>
  	
  	<body>
  		<div class="container">
			<div class="row clearfix">
				<div class="col-md-12 column">
					<div class="row clearfix">
						<div class="col-md-8 column">
							<div class="row clearfix">
								<div class="col-md-12 column">
									<a href="http://blog.cyeam.com" class="thumbnail">
										<img src="http://api.cyeam.com/qr?url={{.ShortUrl}}" alt="Cyeam">
									</a>
								</div>
							</div>
							<div class="row clearfix" style="padding: 30px 0 30px 0;">
								<div class="col-md-12 column">
									<form action="/short_url/shorten" method="post" role="search">
						       				<div class="input-group input-group-lg">
											<input type="text" class="form-control" name="longurl">
											<span class="input-group-btn">
					        					<button class="btn btn-default" type="submit">压缩一下!</button>
					      					</span>
										</div>
						      			</form>
								</div>
							</div>
							<div class="row clearfix">
								<div class="col-md-12 column">
									<form action="/short_url/lengthen" method="post" role="search">
						       				<div class="input-group input-group-lg">
											<input type="text" class="form-control" name="shorturl">
											<span class="input-group-btn">
					        					<button class="btn btn-default" type="submit">还原网址!</button>
					      					</span>
										</div>
						      			</form>
								</div>
							</div>
						</div>
						<div class="col-md-4 column">
							<div class="jumbotron well">
								<div class="alert alert-info" role="alert">
									<a href="{{.ShortUrl}}" target="_blank">{{.ShortUrl}}</a>
								</div>
								<p>
									<a href="{{.LongUrl}}" target="_blank" style="word-wrap:break-word; word-break:break-all">{{.LongUrl}}</a>
								</p>
								<p>

								</p>
							</div>
							<div class="row clearfix" style="padding: 30px 0 30px 0;">
								<div class="col-md-12 column">
									<iframe src="{{.LongUrl}}" class="form-control" style="height: 400px" />
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>

		<script type="text/javascript" src="http://lib.sinaapp.com/js/jquery/1.7.2/jquery.min.js"></script>
		<script src="//maxcdn.bootstrapcdn.com/bootstrap/3.2.0/js/bootstrap.min.js"></script>
	</body>
</html>
