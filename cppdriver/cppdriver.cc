#include <iostream>
#include <memory>
#include <string>
#include <unistd.h>

#include <grpcpp/grpcpp.h>
#include "driver.grpc.pb.h"

#define DRIVER_ID 1


using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
using grpc::ClientReaderWriter;
using controler::Request;
using controler::Response;
using controler::Control;

class DriverClient {
public:
    DriverClient(std::shared_ptr<Channel> channel)
            : stub_(Control::NewStub(channel)) {}

    int streamcall();

private:
    std::unique_ptr<Control::Stub> stub_;
};

int DriverClient::streamcall()  {
    ClientContext context;

    std::shared_ptr<ClientReaderWriter<Request, Response> > stream(
            stub_->StreamCall(&context));

    int seq = 0;
    // login
    {
        Request req;
        req.set_driver_id(DRIVER_ID);
        req.set_seq(seq);
        req.set_data("client login");
        if( !stream->Write(req)) {
            std::cout << "client send login request failed!" << std::endl;
            return 10001;
        }
    }

    Response res;
    while (stream->Read(&res)) {
        std::cout << "Got message " << res.data() << std::endl;

        if (res.ping() == "PING") {
            std::cout << "StreamCall Recevie PING From Server" << std::endl;
            continue;
        }

        if (seq != res.seq()) {
            std::cout << "Seq " << seq << " != " << res.seq() << std::endl;
        }

        // test function, remove it.
        sleep(1);

        Request req;
        req.set_driver_id(DRIVER_ID);
        req.set_seq(seq);
        req.set_data("Action");
        if( !stream->Write(req) ) {
            std::cout << "client send login request failed!" << std::endl;
            return 10002;
        }

        seq++;
    }

    Status status = stream->Finish();
    if (!status.ok()) {
        std::cout << "Stream rpc failed." << std::endl;
    }
}


int main(int argc, char** argv) {
    int cnt = 0;
    while(1) {
        std::cout << "Strat New Connection: " << cnt << std::endl;
        DriverClient driver(grpc::CreateChannel("localhost:8000", grpc::InsecureChannelCredentials()));
        int ret = driver.streamcall();
        std::cout << ret << std::endl;
        sleep(3);
        cnt++;
    }

    return 0;

}
