import { Button } from "@chakra-ui/button";
import { Box, Flex, Text } from "@chakra-ui/layout";
import { useCallback, VFC } from "react";
import { useHistory } from "react-router";

export const TodoCard: VFC =() => {
  const history = useHistory()
  const onClickEditTodo = useCallback(()=> history.push("/todo/edit"),[history])

  return(
    <Box 
      bg="telegram.100"
      borderRadius="md"
      shadow="md"
    >
      <Box p="4"> 
        <Text
          fontSize="2xl"
        >
          これはサンプルです!
        </Text>
      </Box>
      <Box pb="4">
        <Flex justify="space-between">
          <Box>
            <Text
              pl="4"
              fontWeight="bold"
            >
              2021/11/11
            </Text>
          </Box>
          <Box pr="4">
            <Button mr="4" onClick={onClickEditTodo}>編集</Button>
            <Button>削除</Button>
          </Box>
        </Flex>
      </Box>
    </Box>
  )
}