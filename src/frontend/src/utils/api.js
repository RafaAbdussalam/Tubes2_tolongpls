
export async function fetchRecipe({ element, algorithm, mode, maxPaths }) {
    const params = new URLSearchParams({
        element,
        algorithm,
        mode,
        ...(maxPaths && { max_paths: maxPaths })
    });

    try {
        const response = await fetch(`http://localhost:8080/api/recipes?${params}`);
        if (!response.ok) throw new Error('API request failed');
        return await response.json();
    } catch (error) {
        throw error;
    }
};