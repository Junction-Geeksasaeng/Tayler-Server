/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "paper.paper";

export interface MsgCreatePaper {
  creator: string;
  host: string;
  paperName: string;
  owner: string;
  price: string;
}

export interface MsgCreatePaperResponse {
  id: number;
}

export interface MsgUpdatePaper {
  creator: string;
  id: string;
  newOwner: string;
  newPrice: string;
}

export interface MsgUpdatePaperResponse {}

const baseMsgCreatePaper: object = {
  creator: "",
  host: "",
  paperName: "",
  owner: "",
  price: "",
};

export const MsgCreatePaper = {
  encode(message: MsgCreatePaper, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.host !== "") {
      writer.uint32(18).string(message.host);
    }
    if (message.paperName !== "") {
      writer.uint32(26).string(message.paperName);
    }
    if (message.owner !== "") {
      writer.uint32(34).string(message.owner);
    }
    if (message.price !== "") {
      writer.uint32(42).string(message.price);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreatePaper {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreatePaper } as MsgCreatePaper;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.host = reader.string();
          break;
        case 3:
          message.paperName = reader.string();
          break;
        case 4:
          message.owner = reader.string();
          break;
        case 5:
          message.price = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreatePaper {
    const message = { ...baseMsgCreatePaper } as MsgCreatePaper;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.host !== undefined && object.host !== null) {
      message.host = String(object.host);
    } else {
      message.host = "";
    }
    if (object.paperName !== undefined && object.paperName !== null) {
      message.paperName = String(object.paperName);
    } else {
      message.paperName = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = String(object.price);
    } else {
      message.price = "";
    }
    return message;
  },

  toJSON(message: MsgCreatePaper): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.host !== undefined && (obj.host = message.host);
    message.paperName !== undefined && (obj.paperName = message.paperName);
    message.owner !== undefined && (obj.owner = message.owner);
    message.price !== undefined && (obj.price = message.price);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreatePaper>): MsgCreatePaper {
    const message = { ...baseMsgCreatePaper } as MsgCreatePaper;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.host !== undefined && object.host !== null) {
      message.host = object.host;
    } else {
      message.host = "";
    }
    if (object.paperName !== undefined && object.paperName !== null) {
      message.paperName = object.paperName;
    } else {
      message.paperName = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price;
    } else {
      message.price = "";
    }
    return message;
  },
};

const baseMsgCreatePaperResponse: object = { id: 0 };

export const MsgCreatePaperResponse = {
  encode(
    message: MsgCreatePaperResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreatePaperResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreatePaperResponse } as MsgCreatePaperResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreatePaperResponse {
    const message = { ...baseMsgCreatePaperResponse } as MsgCreatePaperResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreatePaperResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreatePaperResponse>
  ): MsgCreatePaperResponse {
    const message = { ...baseMsgCreatePaperResponse } as MsgCreatePaperResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgUpdatePaper: object = {
  creator: "",
  id: "",
  newOwner: "",
  newPrice: "",
};

export const MsgUpdatePaper = {
  encode(message: MsgUpdatePaper, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    if (message.newOwner !== "") {
      writer.uint32(26).string(message.newOwner);
    }
    if (message.newPrice !== "") {
      writer.uint32(34).string(message.newPrice);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdatePaper {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdatePaper } as MsgUpdatePaper;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = reader.string();
          break;
        case 3:
          message.newOwner = reader.string();
          break;
        case 4:
          message.newPrice = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdatePaper {
    const message = { ...baseMsgUpdatePaper } as MsgUpdatePaper;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.newOwner !== undefined && object.newOwner !== null) {
      message.newOwner = String(object.newOwner);
    } else {
      message.newOwner = "";
    }
    if (object.newPrice !== undefined && object.newPrice !== null) {
      message.newPrice = String(object.newPrice);
    } else {
      message.newPrice = "";
    }
    return message;
  },

  toJSON(message: MsgUpdatePaper): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.newOwner !== undefined && (obj.newOwner = message.newOwner);
    message.newPrice !== undefined && (obj.newPrice = message.newPrice);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdatePaper>): MsgUpdatePaper {
    const message = { ...baseMsgUpdatePaper } as MsgUpdatePaper;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.newOwner !== undefined && object.newOwner !== null) {
      message.newOwner = object.newOwner;
    } else {
      message.newOwner = "";
    }
    if (object.newPrice !== undefined && object.newPrice !== null) {
      message.newPrice = object.newPrice;
    } else {
      message.newPrice = "";
    }
    return message;
  },
};

const baseMsgUpdatePaperResponse: object = {};

export const MsgUpdatePaperResponse = {
  encode(_: MsgUpdatePaperResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdatePaperResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdatePaperResponse } as MsgUpdatePaperResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdatePaperResponse {
    const message = { ...baseMsgUpdatePaperResponse } as MsgUpdatePaperResponse;
    return message;
  },

  toJSON(_: MsgUpdatePaperResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgUpdatePaperResponse>): MsgUpdatePaperResponse {
    const message = { ...baseMsgUpdatePaperResponse } as MsgUpdatePaperResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreatePaper(request: MsgCreatePaper): Promise<MsgCreatePaperResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  UpdatePaper(request: MsgUpdatePaper): Promise<MsgUpdatePaperResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreatePaper(request: MsgCreatePaper): Promise<MsgCreatePaperResponse> {
    const data = MsgCreatePaper.encode(request).finish();
    const promise = this.rpc.request("paper.paper.Msg", "CreatePaper", data);
    return promise.then((data) =>
      MsgCreatePaperResponse.decode(new Reader(data))
    );
  }

  UpdatePaper(request: MsgUpdatePaper): Promise<MsgUpdatePaperResponse> {
    const data = MsgUpdatePaper.encode(request).finish();
    const promise = this.rpc.request("paper.paper.Msg", "UpdatePaper", data);
    return promise.then((data) =>
      MsgUpdatePaperResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
