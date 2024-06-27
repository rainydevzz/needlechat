import nacl from "tweetnacl";

export const decrypt = (message: string, nonce: string, pub: string, priv: string) => {
    let shared = nacl.box.before(new Uint8Array(JSON.parse(pub)), new Uint8Array(JSON.parse(priv)))
    let decrypted = nacl.box.open.after(new Uint8Array(JSON.parse(message)), new Uint8Array(JSON.parse(nonce)), shared)
    return decrypted
}