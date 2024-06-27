<script lang="ts">
    import { goto } from "$app/navigation";
    import { request } from "$lib/request";
    import nacl from "tweetnacl";

    let username: string

    const onSubmit = async () => {
        let kp = nacl.box.keyPair()
        let priv = kp.secretKey
        let pub = kp.publicKey
        const r = await request('/user', 'POST', { name: username, publickey: JSON.stringify(Array.from(pub)) })
        localStorage.setItem('publicKey', JSON.stringify(Array.from(pub)) as string)
        localStorage.setItem('privateKey', JSON.stringify(Array.from(priv)) as string)
        localStorage.setItem('username', username)
        localStorage.setItem('id', r.id)
        goto('/')
    }
</script>
<form on:submit|preventDefault={onSubmit}>
    <label>
        username
        <input maxlength="30" minlength="3" required bind:value={username}>
    </label>
    <input type="submit" value="Log in" disabled={!username} />
</form>