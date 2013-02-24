from CGIHTTPServer import CGIHTTPRequestHandler
from BaseHTTPServer import HTTPServer
from cgi import parse_qs
from json import loads
from os import system;

KEYWORD = "RESTART";
HOOK_COMMAND = "echo haha >> /tmp/com";
LOG_FILE = "my_log";


LFILE = open(LOG_FILE, "a+");

class PostHandler(CGIHTTPRequestHandler):
	def do_POST(self):
		try:
			length = int(self.headers.getheader("content-length"));
		except ValueError:
			self.send_response(200);
		post_var = parse_qs(self.rfile.read(length), keep_blank_values=1);
		self.send_response(200);

		data_dict = loads(post_var['payload'][0]);
		commit_list = data_dict['commits'];

		for commit in commit_list:
			if commit['message'].startswith(KEYWORD):
				system(HOOK_COMMAND);
				print commit['message'], "------****", KEYWORD;
				LFILE.write(commit['message']);
				LFILE.write('\n');
				break;

#		print data_dict;

	def do_GET(self):
		self.send_response(200);
		self.send_header("Welcome", "Contact");
		self.end_headers();
		self.wfile.write(self.path);



def run():
	server_address = ('', 9999);
	post_server = HTTPServer(server_address, PostHandler);
	post_server.serve_forever();

if __name__ == '__main__':
	run();
