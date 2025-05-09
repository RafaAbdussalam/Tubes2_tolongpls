export async function fetchRecipe({ element, algorithm, maxPaths }) {
    // Konversi parameter frontend ke format backend
    const backendParams = {
        element: element,
        mode: algorithm,    // 'algorithm' di frontend -> 'mode' di backend
        amount: maxPaths    // Konversi mode frontend ke amount
    };

    const params = new URLSearchParams(backendParams);

    try {
        const response = await fetch(
            `http://localhost:8080/api/tree?${params}`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json'
                }
            }
        );
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        } else {
            const contentType = response.headers.get('content-type');
            if (!contentType?.includes('application/json')) {
                throw new TypeError("Response is not JSON");
            }
        }
        return await response.json();
    } catch (error) {
        alert(`
            - error: ${error}
            - element: ${backendParams.element}
            - mode: ${backendParams.mode}
            - amount: ${backendParams.amount}
        `);
        throw error;
    }
}