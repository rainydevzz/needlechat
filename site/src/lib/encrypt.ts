import nacl from "tweetnacl";

export const encrypt = (pub: string, priv: string, content: string): {message: Uint8Array, nonce: Uint8Array} => {
    let t = new TextEncoder()
    let n = nacl.randomBytes(24)
    let shared = nacl.box.before(new Uint8Array(JSON.parse(pub)), new Uint8Array(JSON.parse(priv)))
    let encrypted = nacl.box.after(t.encode(content), n, shared)
    return {message: encrypted, nonce: n}
}