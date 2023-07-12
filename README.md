<div class="header" align="center">
  <p>
    <img src="https://static.wikia.nocookie.net/pokemonwack/images/a/a3/903.png/revision/latest?cb=20191207232414" />
  </p>
  <h2>klefki (a bitwarden-inspired password vault)</h2>
</div>

<div class="requirements">
  <h4>requirements:</h4>
  <ul>
    <li>Golang</li>
      <ul>
        <li>air</li>
        <li>sqlc</li>
      </ul>
    <li>Node</li>
    <ul>
      <li>Svelte</li>
      <li>Tailwind</li>
    </ul>
  </ul>
</div> 

### Getting started

Install dependencies:

```
make install
```

Start frontend/backend independently:

```
# backend
make start-backend

# frontend
make start-frontend
```

Serve static files with backend:
```
make preview
```