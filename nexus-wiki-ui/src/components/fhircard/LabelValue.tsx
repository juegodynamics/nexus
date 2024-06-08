import { Box, Stack, Typography, useTheme } from "@mui/material";
import { FadeBlock } from "../ux";

export const LabelValue = ({
  label,
  value,
}: {
  label: string;
  value: string;
}) => {
  const theme = useTheme();
  return (
    <Stack spacing={0} sx={{ pt: 0, pb: 0, mt: 0, mb: 0 }}>
      <Typography
        variant="caption"
        sx={{
          color: theme.palette.primary.main,
          pb: -0.5,
          mb: -0.5,
          pt: -0.5,
          mt: -0.5,
        }}
      >
        {label}
      </Typography>
      <Typography
        variant="body2"
        sx={{
          color: theme.palette.grey[100],
          pt: 0,
          mt: 0,
          mb: 1,
        }}
      >
        {value}
      </Typography>
    </Stack>
  );
};
