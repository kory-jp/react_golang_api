import { VFC } from "react"
import { Box, Text } from "@chakra-ui/react"

export const Top: VFC = () => {
  return(
    <Box w="100%">
      <Box h="10em" position="relative">
        <Text
          m="0"
          position="absolute"
          top="50%"
          left="40%"
          fontSize="30px"
          fontWeight="bold"
        >
          TodoAppへようこそ!
        </Text>
      </Box>
    </Box>
  )
}

export default Top;