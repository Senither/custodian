// @ts-check
import antfu from '@antfu/eslint-config'

export default antfu({
    ignores: ['lib/**.ts', 'docker-compose.yml'],
    stylistic: {
        indent: 4,
    },
})
