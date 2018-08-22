using System;
using System.Reflection;
using System.Runtime.InteropServices.ComTypes;
using Improbable;
using Improbable.Collections;
using Improbable.Worker;

namespace Managed
{
    internal class Startup
    {
        private const string WorkerType = "Managed";

        private const string LoggerName = "Startup.cs";

        private const int ErrorExitStatus = 1;

        private const uint GetOpListTimeoutInMilliseconds = 100;


        private static WorkerRequirementSet MakeRequirements(string value)
        {
            var attributeSet = new WorkerAttributeSet(new List<string>{value});
            return new WorkerRequirementSet(new List<WorkerAttributeSet>{attributeSet});
        }
        
        private static Entity MakeTree(double x, double y, double z)
        {
            var entity = new Entity();
            entity.Add(new Metadata.Data("Tree"));
            entity.Add(new Position.Data(new Coordinates(x,y,z)));
            entity.Add(new EntityAcl.Data(MakeRequirements("player"), new Map<uint, WorkerRequirementSet>{}));
            entity.Add(new Persistence.Data());
            return entity;
        }
        
        private static Entity MakePlayer()
        {
            var entity = new Entity();
            entity.Add(new Metadata.Data("Player"));
            entity.Add(new Position.Data(new Coordinates(0,0,0)));
            
            var write = new Map<uint, WorkerRequirementSet>();
            write[Position.ComponentId] = MakeRequirements("player");
            
            entity.Add(new EntityAcl.Data(MakeRequirements("player"), write));
            entity.Add(new Persistence.Data());
            return entity;
        }
        
        private static int Main(string[] args)
        {
            var entities = new Map<EntityId, Entity>();
            var id = 1L;
            entities[new EntityId(id++)] = MakeTree(1,2,3);
            entities[new EntityId(id++)] = MakePlayer();

            var output = Snapshot.Save("D:\\Projects\\CsharpBlankProject\\snapshots\\default.snapshot", entities);
            
            Console.WriteLine(output.Value);
            
            if (args.Length != 4) {
                PrintUsage();
                return ErrorExitStatus;
            }

            // Avoid missing component errors because no components are directly used in this project
            // and the GeneratedCode assembly is not loaded but it should be
            Assembly.Load("GeneratedCode");

            var connectionParameters = new ConnectionParameters
            {
                WorkerType = WorkerType,
                Network =
                {
                    ConnectionType = NetworkConnectionType.Tcp
                }
            };

            using (var connection = ConnectWithReceptionist(args[1], Convert.ToUInt16(args[2]), args[3], connectionParameters))
            using (var dispatcher = new Dispatcher())
            {
                var isConnected = true;

                dispatcher.OnDisconnect(op =>
                {
                    Console.Error.WriteLine("[disconnect] " + op.Reason);
                    isConnected = false;
                });

                dispatcher.OnLogMessage(op =>
                {
                    connection.SendLogMessage(op.Level, LoggerName, op.Message);
                    if (op.Level == LogLevel.Fatal)
                    {
                        Console.Error.WriteLine("Fatal error: " + op.Message);
                        Environment.Exit(ErrorExitStatus);
                    }
                });

                while (isConnected)
                {
                    using (var opList = connection.GetOpList(GetOpListTimeoutInMilliseconds))
                    {
                        dispatcher.Process(opList);
                    }
                }
            }

            // This means we forcefully disconnected
            return ErrorExitStatus;
        }

        private static void PrintUsage()
        {
            Console.WriteLine("Usage: mono Managed.exe receptionist <hostname> <port> <worker_id>");
            Console.WriteLine("Connects to SpatialOS");
            Console.WriteLine("    <hostname>      - hostname of the receptionist to connect to.");
            Console.WriteLine("    <port>          - port to use");
            Console.WriteLine("    <worker_id>     - name of the worker assigned by SpatialOS.");
        }

        private static Connection ConnectWithReceptionist(string hostname, ushort port,
            string workerId, ConnectionParameters connectionParameters)
        {
            Connection connection;

            // You might want to change this to true or expose it as a command-line option
            // if using `spatial cloud connect external` for debugging
            connectionParameters.Network.UseExternalIp = false;

            using (var future = Connection.ConnectAsync(hostname, port, workerId, connectionParameters))
            {
                connection = future.Get();
            }

            connection.SendLogMessage(LogLevel.Info, LoggerName, "Successfully connected using the Receptionist");

            return connection;
        }
    }
}