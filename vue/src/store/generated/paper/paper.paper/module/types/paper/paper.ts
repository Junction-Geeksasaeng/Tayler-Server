/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "paper.paper";

export interface Paper {
  creator: string;
  id: number;
  host: string;
  paperName: string;
  owner: string;
  price: string;
}

const basePaper: object = {
  creator: "",
  id: 0,
  host: "",
  paperName: "",
  owner: "",
  price: "",
};

export const Paper = {
  encode(message: Paper, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.host !== "") {
      writer.uint32(26).string(message.host);
    }
    if (message.paperName !== "") {
      writer.uint32(34).string(message.paperName);
    }
    if (message.owner !== "") {
      writer.uint32(42).string(message.owner);
    }
    if (message.price !== "") {
      writer.uint32(50).string(message.price);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Paper {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePaper } as Paper;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.host = reader.string();
          break;
        case 4:
          message.paperName = reader.string();
          break;
        case 5:
          message.owner = reader.string();
          break;
        case 6:
          message.price = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Paper {
    const message = { ...basePaper } as Paper;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
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

  toJSON(message: Paper): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.host !== undefined && (obj.host = message.host);
    message.paperName !== undefined && (obj.paperName = message.paperName);
    message.owner !== undefined && (obj.owner = message.owner);
    message.price !== undefined && (obj.price = message.price);
    return obj;
  },

  fromPartial(object: DeepPartial<Paper>): Paper {
    const message = { ...basePaper } as Paper;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
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
