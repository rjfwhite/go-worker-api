

Had to update `c_worker.h`
```
//WORKER_API Schema_FieldId
//Schema_GetCommandResponseComponentId(const Schema_CommandResponse* request);
/** Get the 1-based position of the command in the order the commands appear in the schema. */
```


Had to remove 

```
/** Acquire a reference to extend the lifetime of a command request owned by the SDK. */
Worker_CommandRequest* Worker_AcquireCommandRequest(const Worker_CommandRequest* request);
/** Acquire a reference to extend the lifetime of a command response owned by the SDK. */
Worker_CommandResponse* Worker_AcquireCommandResponse(const Worker_CommandResponse* response);
/** Acquire a reference to extend the lifetime of a component data snapshot owned by the SDK. */
Worker_ComponentData* Worker_AcquireComponentData(const Worker_ComponentData* data);
/** Acquire a reference to extend the lifetime of a component update owned by the SDK. */
Worker_ComponentUpdate* Worker_AcquireComponentUpdate(const Worker_ComponentUpdate* update);
/** Release a reference obtained by Worker_AcquireCommandRequest. */
void Worker_ReleaseCommandRequest(Worker_CommandRequest* request);
/** Release a reference obtained by Worker_AcquireCommandResponse. */
void Worker_ReleaseCommandResponse(Worker_CommandResponse* response);
/** Release a reference obtained by Worker_AcquireComponentData. */
void Worker_ReleaseComponentData(Worker_ComponentData* data);
/** Release a reference obtained by Worker_AcquireComponentUpdate. */
void Worker_ReleaseComponentUpdate(Worker_ComponentUpdate* update);
```

Shennanigans around having to copy `SwigcptrStruct_SS_Worker_OpList` as opposed to `SwigcptrWorker_Oplist` etc - seems like a bug in the golang swig generator