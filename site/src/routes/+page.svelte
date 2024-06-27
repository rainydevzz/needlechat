<script lang="ts">
    import { decrypt } from '$lib/decrypt';
    import { encrypt } from '$lib/encrypt';
    import { request } from '$lib/request'
    import { onMount } from 'svelte'

    let arr: Array<any> = []
    let content: string

    export const load = async () => {
        let t = []
        let r = await request('/', 'GET')
        let k = await request('/users', 'GET')
        for(const i of k) {
            if(i.PublicKey != localStorage.getItem('publicKey') as string) {
                k = i;
                break;
            }
        }
        for(const i of r) {
            let d = (decrypt(i.Content, i.Nonce, k.PublicKey, localStorage.getItem('privateKey') as string) as Uint8Array)
            t.push({message: String.fromCharCode(...d), name: k.Username})
        }
        arr = t
    }

    onMount(async () => {
        await load()
    })

    const onSubmit = async () => {
        let k = await request('/users', 'GET')
        for(const i of k) {
            if(i.PublicKey != localStorage.getItem('publicKey') as string) {
                k = i.PublicKey;
                break;
            }
        }
        let e = encrypt(k, localStorage.getItem('privateKey') as string, content)
        let id = localStorage.getItem('username') as string
        await request('/message', 'POST', {content: JSON.stringify(Array.from(e.message)), username: id, nonce: JSON.stringify(Array.from(e.nonce))})
        window.location.href = '/'
    }

</script>
{#each arr as a}
    <h1>{a.message} - {a.name}</h1>
{/each}
<form on:submit|preventDefault={onSubmit}>
    <label>
        message
        <input maxlength="30" minlength="3" required bind:value={content}>
    </label>
    <input type="submit" value="Send" disabled={!content} />
</form>