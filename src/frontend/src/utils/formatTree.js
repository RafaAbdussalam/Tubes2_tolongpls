// Helper function untuk membuat formatted tree agar mudah divisualisasikan di RecipeTree.js

export function formatTree(rawData) {
    function transformNode(node) {
        if (!node.children || node.children.length === 0) {
            return {
                name: node.element,
                attributes: { 
                    type: 'base-element' 
                }
            };
        }

        // Proses setiap resep (pasangan item)
        return {
            name: node.element,
            children: node.children.map(recipe => ({
                name: `${recipe.item_1.element} + ${recipe.item_2.element}`,
                attributes: { 
                    type: 'recipe' 
                },
                children: [
                    transformNode(recipe.item_1),
                    transformNode(recipe.item_2)
                ]
            }))
        };
    };

    return transformNode(rawData);
};