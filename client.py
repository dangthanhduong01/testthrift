import sys
sys.path.append('./gen-py')

from trythrift import Calculator
from trythrift.ttypes import *

from thrift import Thrift
from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol

def main():
    try:
        transport = TSocket.TSocket('localhost', 9090)
        transport = TTransport.TBufferedTransport(transport)
        protocol = TBinaryProtocol.TBinaryProtocol(transport)
        client = Calculator.Client(protocol)

        transport.open()

        num1 = 10
        num2 = 5

        result_add = client.add(num1, num2)
        print(f"{num1} + {num2} = {result_add}")

        result_multiply = client.multiply(num1, num2)
        print(f"{num1} * {num2} = {result_multiply}")

        transport.close()

    except Thrift.TException as tx:
        print(f"Thrift Exception: {tx.message}")

if __name__ == "__main__":
    main()
