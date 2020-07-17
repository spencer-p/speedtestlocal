result = document.getElementById('result');
document.getElementById('speedtest').addEventListener('click', start_test);

var tot_bytes = 0;
var tStart = undefined;

function start_test() {
	// create ws connection
	const host = window.location.host;
	const socket = new WebSocket('ws://'+host+'/ws');

	// initialize test when the conn opens
	socket.addEventListener('open', function() {
		console.log('connection opened');
		tot_bytes = 0;
		tStart = Date.now();
	});

	// listen for messages and recompute speed
	socket.addEventListener('message', function(event) {
		const nbytes = event.data.length;
		compute_speed(nbytes);
	});

	// on close finalize the speed
	socket.addEventListener('close', function(event) {
		console.log('closed');
		comptue_speed(0);
	});
}

function compute_speed(nbytes) {
	tot_bytes += nbytes;
	const tot_time = Date.now() - tStart;
	const Mb = tot_bytes/(1024*1024);
	const sec = tot_time/1000;
	const Mbps = Mb/sec;
	result.innerText = `${Mbps.toFixed(2)} Mbps (${Mb} mb in ${sec} seconds)`
}
